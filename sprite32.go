package gosang

// Sprite32 is a 32-bit color sprite.
type Sprite32 struct {
	width  int
	height int
	count  int
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
