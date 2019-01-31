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
