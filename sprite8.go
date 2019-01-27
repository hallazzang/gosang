package gosang

import "image"

// Sprite8 is an 8-bit color sprite.
type Sprite8 struct {
	width  int
	height int
	count  int
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

func (sp *Sprite8) Frame(idx int) (image.Image, error) {
	return nil, nil
}
