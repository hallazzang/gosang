package gosang

import "image"

// Sprite32 is a 32-bit color sprite.
type Sprite32 struct {
	r      reader
	width  int
	height int
	count  int
}

func newSprite32(r reader, header spriteHeader) (*Sprite32, error) {
	sp := &Sprite32{
		r:      r,
		width:  int(header.Width),
		height: int(header.Height),
		count:  int(header.Count),
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
	return nil, nil
}
