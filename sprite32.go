package gosang

import (
	"encoding/binary"
	"image"

	"github.com/pkg/errors"
)

// Sprite32 is a 32-bit color sprite.
type Sprite32 struct {
	r       reader
	width   int
	height  int
	count   int
	offsets []uint32
}

func newSprite32(r reader, header spriteHeader) (*Sprite32, error) {
	sp := &Sprite32{
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

// ColorBits returns sprite's color bits. This method always returns 32
// for Sprite32.
func (sp *Sprite32) ColorBits() int {
	return 32
}

// Width returns sprite's frame width, in pixel.
func (sp *Sprite32) Width() int {
	return sp.width
}

// Height returns sprite's frame height, in pixel.
func (sp *Sprite32) Height() int {
	return sp.height
}

// Count returns sprite's frame count.
func (sp *Sprite32) Count() int {
	return sp.count
}

// Frame returns specific frame's data as image.Image.
func (sp *Sprite32) Frame(idx int) (image.Image, error) {
	if idx < 0 || idx > sp.count-1 {
		return nil, errors.New("invalid frame index")
	}
	return nil, nil
}
