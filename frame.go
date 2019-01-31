package gosang

import (
	"image"
)

// Frame represents single frame in sprite.
type Frame struct {
	sp     Sprite
	idx    int
	width  int
	height int
	img    image.Image
}

func newFrame(sp Sprite, idx int, img image.Image) *Frame {
	return &Frame{sp, idx, sp.FrameWidth(), sp.FrameHeight(), img}
}

// Index returns frame's index in sprite.
func (fr *Frame) Index() int {
	return fr.idx
}

// Width returns frame's width, in pixels.
func (fr *Frame) Width() int {
	return fr.width
}

// Height returns frame's height, in pixels.
func (fr *Frame) Height() int {
	return fr.height
}

// Image returns frame's image.
func (fr *Frame) Image() image.Image {
	return fr.img
}
