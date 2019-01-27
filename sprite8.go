package gosang

// Sprite8 is an 8-bit color sprite.
type Sprite8 struct {
	width  int
	height int
	count  int
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
