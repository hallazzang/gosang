package gosang

import (
	"bufio"
	"encoding/binary"
	"image"
	"image/color"
	"io"

	"github.com/pkg/errors"
)

type sprite32Alpha struct {
	spriteBase
}

func newSprite32Alpha(r io.ReaderAt, header spriteHeader) (*sprite32Alpha, error) {
	sp := &sprite32Alpha{spriteBase{
		r:           r,
		frameWidth:  header.FrameWidth,
		frameHeight: header.FrameHeight,
		frameCount:  header.FrameCount,
		offsets:     make([]uint32, header.FrameCount),
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

func (sp *sprite32Alpha) Frame(idx int) (*Frame, error) {
	if idx < 0 || idx > int(sp.frameCount-1) {
		return nil, errors.New("frame index out of range")
	}
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
	return newFrame(sp, idx, img), nil
}

type sprite32AlphaPixel struct{ Alpha, Red, Green, Blue uint8 }
