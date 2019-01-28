package gosang

import (
	"bufio"
	"encoding/binary"
	"image"
	"io"

	"github.com/pkg/errors"
)

// sprite8 is an 8-bit color sprite.
type sprite8 struct {
	r           io.ReaderAt
	frameWidth  uint32
	frameHeight uint32
	frameCount  uint32
	offsets     []uint32
	width       uint32
	height      uint32
}

func newSprite8(r io.ReaderAt, header spriteHeader) (*sprite8, error) {
	sp := &sprite8{
		r:           r,
		frameWidth:  header.FrameWidth,
		frameHeight: header.FrameHeight,
		frameCount:  header.FrameCount,
		offsets:     make([]uint32, header.FrameCount),
	}
	if err := binary.Read(&offsetedReader{r, 0x4c0}, binary.LittleEndian, &sp.offsets); err != nil {
		return nil, errors.Wrap(err, "failed to read frame offsets")
	}
	if err := binary.Read(&offsetedReader{r, 0xbcc}, binary.LittleEndian, &sp.width); err != nil {
		return nil, errors.Wrap(err, "failed to read sprite width")
	}
	if err := binary.Read(&offsetedReader{r, 0xbd0}, binary.LittleEndian, &sp.height); err != nil {
		return nil, errors.Wrap(err, "failed to read sprite height")
	}
	return sp, nil
}

func (sp *sprite8) ColorBits() int {
	return 8
}

func (sp *sprite8) FrameWidth() int {
	return int(sp.frameWidth)
}

func (sp *sprite8) FrameHeight() int {
	return int(sp.frameHeight)
}

func (sp *sprite8) FrameCount() int {
	return int(sp.frameCount)
}

func (sp *sprite8) Width() int {
	return int(sp.width)
}

func (sp *sprite8) Height() int {
	return int(sp.height)
}

func (sp *sprite8) Frame(idx int) (image.Image, error) {
	if idx < 0 || idx > int(sp.frameCount-1) {
		return nil, errors.New("frame index out of range")
	}
	img := image.NewPaletted(image.Rect(0, 0, int(sp.frameWidth), int(sp.frameHeight)), sprite8Palette)
	r := bufio.NewReader(&offsetedReader{sp.r, 0xbf4 + int64(sp.offsets[idx])})
	for y := uint32(0); y < sp.frameHeight; y++ {
		for x := uint32(0); x < sp.frameWidth; {
			b, err := r.ReadByte()
			if err != nil {
				return nil, errors.Wrap(err, "failed to read frame data")
			}
			if b == 0xfe {
				c, err := r.ReadByte()
				if err != nil {
					return nil, errors.Wrap(err, "failed to read frame data")
				}
				for ; c > 0; c-- {
					img.SetColorIndex(int(x), int(y), b)
					x++
				}
			} else {
				img.SetColorIndex(int(x), int(y), b)
				x++
			}
		}
	}
	return img, nil
}
