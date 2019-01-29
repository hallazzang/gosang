package gosang

import (
	"bufio"
	"encoding/binary"
	"image"
	"image/color"
	"io"

	"github.com/pkg/errors"
)

// sprite32 is a 32-bit color sprite.
type sprite32 struct {
	r           io.ReaderAt
	frameWidth  uint32
	frameHeight uint32
	frameCount  uint32
	offsets     []uint32
	width       uint32
	height      uint32
}

func newSprite32(r io.ReaderAt, header spriteHeader) (*sprite32, error) {
	sp := &sprite32{
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

func (sp *sprite32) ColorBits() int {
	return 32
}

func (sp *sprite32) FrameWidth() int {
	return int(sp.frameWidth)
}

func (sp *sprite32) FrameHeight() int {
	return int(sp.frameHeight)
}

func (sp *sprite32) FrameCount() int {
	return int(sp.frameCount)
}

func (sp *sprite32) Width() int {
	return int(sp.width)
}

func (sp *sprite32) Height() int {
	return int(sp.height)
}

func (sp *sprite32) Frame(idx int) (image.Image, error) {
	if idx < 0 || idx > int(sp.frameCount-1) {
		return nil, errors.New("frame index out of range")
	}
	img := image.NewRGBA(image.Rect(0, 0, int(sp.frameWidth), int(sp.frameHeight)))
	r := bufio.NewReader(&offsetedReader{sp.r, 0xe4c + int64(sp.offsets[idx])})
	for y := 0; y < int(sp.frameHeight); y++ {
		for x := 0; x < int(sp.frameWidth); {
			var p struct{ Count, Blue, Green, Red byte }
			if err := binary.Read(r, binary.LittleEndian, &p); err != nil {
				return nil, errors.Wrap(err, "failed to read frame data")
			}
			for ; p.Count > 0; p.Count-- {
				img.SetRGBA(int(x), int(y), color.RGBA{p.Red, p.Green, p.Blue, 0xff})
				x++
			}
		}
	}
	return img, nil
}

func (sp *sprite32) frameOffset(idx int) (int64, error) {
	return 0, nil
}

func (sp *sprite32) frameSize(idx int) (int, error) {
	return 0, nil
}
