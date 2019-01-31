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
	r           io.ReaderAt
	frameWidth  uint32
	frameHeight uint32
	frameCount  uint32
	offsets     []uint32
	width       uint32
	height      uint32
	lastOffset  uint32
	frames      []*Frame
}

func newSprite32Alpha(r io.ReaderAt, header spriteHeader) (*sprite32Alpha, error) {
	sp := &sprite32Alpha{
		r:           r,
		frameWidth:  header.FrameWidth,
		frameHeight: header.FrameHeight,
		frameCount:  header.FrameCount,
		offsets:     make([]uint32, header.FrameCount),
	}
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

func (sp *sprite32Alpha) FrameWidth() int {
	return int(sp.frameWidth)
}

func (sp *sprite32Alpha) FrameHeight() int {
	return int(sp.frameHeight)
}

func (sp *sprite32Alpha) FrameCount() int {
	return int(sp.frameCount)
}

func (sp *sprite32Alpha) Width() int {
	return int(sp.width)
}

func (sp *sprite32Alpha) Height() int {
	return int(sp.height)
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

func (sp *sprite32Alpha) frameOffset(idx int) (int64, error) {
	if idx < 0 || idx > int(sp.frameCount-1) {
		return 0, errors.New("frame index out of range")
	}
	return int64(sp.offsets[idx]), nil
}

func (sp *sprite32Alpha) frameSize(idx int) (int, error) {
	if idx < 0 || idx > int(sp.frameCount-1) {
		return 0, errors.New("frame index out of range")
	} else if idx < int(sp.frameCount-1) {
		return int(sp.offsets[idx+1] - sp.offsets[idx]), nil
	}
	if sp.lastOffset == 0 {
		if err := binary.Read(&offsetedReader{sp.r, 0xe20}, binary.LittleEndian, &sp.lastOffset); err != nil {
			return 0, errors.Wrap(err, "failed to read sprite's last data offset")
		}
	}
	return int(sp.lastOffset - sp.offsets[idx]), nil
}

type sprite32AlphaPixel struct{ Alpha, Red, Green, Blue uint8 }
