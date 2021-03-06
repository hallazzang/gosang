package gosang

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"image"
	"image/color"
	"io"

	"github.com/pkg/errors"
)

type sprite32Alpha struct {
	sprite
}

func newSprite32Alpha(r io.ReaderAt, header spriteHeader) (*sprite32Alpha, error) {
	sp := &sprite32Alpha{sprite{
		r:           r,
		frameWidth:  header.FrameWidth,
		frameHeight: header.FrameHeight,
		frameCount:  header.FrameCount,
		offsets:     make([]uint32, header.FrameCount),
		frames:      make([]*Frame, header.FrameCount),
	}}
	if err := binary.Read(&offsetedReader{r, 0x4c0}, binary.LittleEndian, &sp.offsets); err != nil {
		return nil, errors.Wrap(err, "failed to read frame offsets")
	}
	if err := binary.Read(&offsetedReader{r, 0xe24}, binary.LittleEndian, &sp.width); err != nil {
		return nil, errors.Wrap(err, "failed to read sprite width")
	}
	if err := binary.Read(&offsetedReader{r, 0xe28}, binary.LittleEndian, &sp.height); err != nil {
		return nil, errors.Wrap(err, "failed to read sprite height")
	}
	return sp, nil
}

func (sp *sprite32Alpha) ColorBits() int {
	return 32
}

func (sp *sprite32Alpha) HasAlpha() bool {
	return true
}

func (sp *sprite32Alpha) Save(w io.Writer) error {
	header := spriteHeader{
		Signature:   0x19,
		FrameWidth:  sp.frameWidth,
		FrameHeight: sp.frameHeight,
		FrameCount:  sp.frameCount,
	}
	if err := binary.Write(w, binary.LittleEndian, &header); err != nil {
		return errors.Wrap(err, "failed to write sprite header")
	}
	offsets := make([]uint32, sp.frameCount)
	sizes := make([]uint32, sp.frameCount)
	if err := advanceWriter(w, 0x4c0-16); err != nil {
		return errors.Wrap(err, "failed to advance writer")
	}
	buf := new(bytes.Buffer)
	offset := uint32(0)
	for i := uint32(0); i < sp.frameCount; i++ {
		if sp.frames[i] == nil {
			return errors.Errorf("frame #%d is empty", i)
		}
		size, err := sp.encodeFrame(buf, int(i))
		if err != nil {
			return errors.Wrapf(err, "failed to encode frame #%d", i)
		}
		offsets[i] = offset
		sizes[i] = uint32(size)
		offset += sizes[i]
	}
	if err := binary.Write(w, binary.LittleEndian, offsets); err != nil {
		return errors.Wrap(err, "failed to write frame offsets")
	}
	if err := advanceWriter(w, int(0x970-(0x4c0+4*sp.frameCount))); err != nil {
		return errors.Wrap(err, "failed to advance writer")
	}
	for i := uint32(0); i < sp.frameCount; i++ {
		if err := binary.Write(w, binary.LittleEndian, sizes[i]/4); err != nil {
			return errors.Wrap(err, "failed to write encoded frame size")
		}
	}
	if err := advanceWriter(w, int(0xe20-(0x970+4*sp.frameCount))); err != nil {
		return errors.Wrap(err, "failed to advance writer")
	}
	if err := binary.Write(w, binary.LittleEndian, offsets[len(offsets)-1]+sizes[len(sizes)-1]); err != nil {
		return errors.Wrap(err, "failed to write frame data size")
	}
	if err := binary.Write(w, binary.LittleEndian, sp.width); err != nil {
		return errors.Wrap(err, "failed to write frame width")
	}
	if err := binary.Write(w, binary.LittleEndian, sp.height); err != nil {
		return errors.Wrap(err, "failed to write frame height")
	}
	if err := advanceWriter(w, 0xe4c-(0xe20+12)); err != nil {
		return errors.Wrap(err, "failed to advance writer")
	}
	if _, err := buf.WriteTo(w); err != nil {
		return errors.Wrap(err, "failed to write frame data")
	}
	return nil
}

func (sp *sprite32Alpha) loadFrame(idx int) (*Frame, error) {
	if idx < 0 || idx > int(sp.frameCount-1) {
		return nil, errors.New("frame index out of range")
	}
	if sp.frames[idx] == nil {
		img := image.NewNRGBA(image.Rect(0, 0, int(sp.frameWidth), int(sp.frameHeight)))
		r := bufio.NewReader(&offsetedReader{sp.r, 0xe4c + int64(sp.offsets[idx])})
		for y := 0; y < int(sp.frameHeight); y++ {
			for x := 0; x < int(sp.frameWidth); {
				var p sprite32AlphaPixel
				if err := binary.Read(r, binary.LittleEndian, &p); err != nil {
					return nil, errors.Wrap(err, "failed to read frame data")
				}
				if p.Alpha == 0 && p.Green == 0 && p.Blue == 0 {
					for ; p.Red > 0; p.Red-- {
						img.SetNRGBA(int(x), int(y), color.NRGBA{0xfc, 0xe0, 0xfc, 0x00})
						x++
					}
				} else {
					img.SetNRGBA(int(x), int(y), color.NRGBA{p.Red, p.Green, p.Blue, p.Alpha})
					x++
				}
			}
		}
		sp.frames[idx] = newFrame(sp, idx, img)
	}
	return sp.frames[idx], nil
}

func (sp *sprite32Alpha) encodeFrame(w io.Writer, idx int) (int, error) {
	if idx < 0 || idx > int(sp.frameCount-1) {
		return 0, errors.New("frame index out of range")
	}
	img := sp.frames[idx].img
	if img == nil {
		return 0, errors.Errorf("frame #%d's image is empty", idx)
	}
	b := img.Bounds()
	if b.Empty() {
		return 0, errors.Errorf("invalid image bounds: %v", b)
	}
	width, height := b.Dx(), b.Dy()
	if uint32(width) != sp.frameWidth || uint32(height) != sp.frameHeight {
		return 0, errors.New("mismatched frame size")
	}
	n := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; {
			var r, g, b, a uint8
			c := uint8(0)
			for x < width && c < 0xff {
				p := img.At(x, y)
				switch p := p.(type) {
				default:
					tr, tg, tb, ta := p.RGBA()
					r, g, b, a = uint8(tr), uint8(tg), uint8(tb), uint8(ta)
				case color.NRGBA:
					r, g, b, a = p.R, p.G, p.B, p.A
				}
				x++
				if a == 0 {
					c++
				} else {
					break
				}
			}
			if c > 0 {
				if err := binary.Write(w, binary.LittleEndian, sprite32AlphaPixel{0, c, 0, 0}); err != nil {
					return n, errors.Wrap(err, "failed to write frame data")
				}
				n += 4
			}
			if a != 0 {
				if err := binary.Write(w, binary.LittleEndian, sprite32AlphaPixel{uint8(a), uint8(r), uint8(g), uint8(b)}); err != nil {
					return n, errors.Wrap(err, "failed to write frame data")
				}
				n += 4
			}
		}
	}
	return n, nil
}

type sprite32AlphaPixel struct{ Alpha, Red, Green, Blue uint8 }
