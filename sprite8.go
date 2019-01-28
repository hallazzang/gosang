package gosang

import (
	"bufio"
	"encoding/binary"
	"image"

	"github.com/pkg/errors"
)

// sprite8 is an 8-bit color sprite.
type sprite8 struct {
	r       reader
	width   int
	height  int
	count   int
	offsets []uint32
}

func newSprite8(r reader, header spriteHeader) (*sprite8, error) {
	sp := &sprite8{
		r:       r,
		width:   int(header.Width),
		height:  int(header.Height),
		count:   int(header.Count),
		offsets: make([]uint32, header.Count),
	}
	if err := binary.Read(&offsetedReader{r, 0x4c0}, binary.LittleEndian, &sp.offsets); err != nil {
		return nil, errors.Wrap(err, "failed to read frame offsets")
	}
	return sp, nil
}

// ColorBits returns sprite's color bits. This method always returns 8
// for sprite8.
func (sp *sprite8) ColorBits() int {
	return 8
}

// Width returns sprite's frame width, in pixel.
func (sp *sprite8) Width() int {
	return sp.width
}

// Height returns sprite's frame height, in pixel.
func (sp *sprite8) Height() int {
	return sp.height
}

// Count returns sprite's frame count.
func (sp *sprite8) Count() int {
	return sp.count
}

// Frame returns specific frame's data as image.Image.
func (sp *sprite8) Frame(idx int) (image.Image, error) {
	if idx < 0 || idx > sp.count-1 {
		return nil, errors.New("frame index out of range")
	}
	img := image.NewPaletted(image.Rect(0, 0, sp.width, sp.height), sprite8Palette)
	r := bufio.NewReader(&offsetedReader{sp.r, 0xbf4 + int64(sp.offsets[idx])})
	for y := 0; y < sp.height; y++ {
		for x := 0; x < sp.width; {
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
					img.SetColorIndex(x, y, b)
					x++
				}
			} else {
				img.SetColorIndex(x, y, b)
				x++
			}
		}
	}
	return img, nil
}
