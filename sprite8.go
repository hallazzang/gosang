package gosang

import (
	"encoding/binary"
	"image"

	"github.com/pkg/errors"
)

// Sprite8 is an 8-bit color sprite.
type Sprite8 struct {
	r       reader
	width   int
	height  int
	count   int
	offsets []uint32
}

func newSprite8(r reader, header spriteHeader) (*Sprite8, error) {
	sp := &Sprite8{
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
// for Sprite8.
func (sp *Sprite8) ColorBits() int {
	return 8
}

// Width returns sprite's frame width, in pixel.
func (sp *Sprite8) Width() int {
	return sp.width
}

// Height returns sprite's frame height, in pixel.
func (sp *Sprite8) Height() int {
	return sp.height
}

// Count returns sprite's frame count.
func (sp *Sprite8) Count() int {
	return sp.count
}

// Frame returns specific frame's data as image.Image.
func (sp *Sprite8) Frame(idx int) (image.Image, error) {
	return nil, nil
}

func (sp *Sprite8) loadOffsets() error {
	sp.offsets = make([]uint32, sp.count)
	if err := binary.Read(sp.r, binary.LittleEndian, &sp.offsets); err != nil {
		return err
	}
	return nil
}
