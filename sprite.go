package gosang

import (
	"encoding/binary"
	"image"
	"io"

	"github.com/pkg/errors"
)

// Sprite represents single sprite. It can either be 8-bit or 32-bit sprite.
type Sprite interface {
	ColorBits() int   // Color bits. 8 or 32.
	FrameWidth() int  // Frame's width in pixels.
	FrameHeight() int // Frame's height in pixels.
	FrameCount() int
	Width() int
	Height() int
	Frame(idx int) (image.Image, error) // Specific frame's data.

	frameOffset(idx int) (int64, error)
	frameSize(idx int) (int, error)
}

// NewSprite creates new sprite from r.
func NewSprite(r io.ReaderAt) (Sprite, error) {
	var header spriteHeader
	if err := binary.Read(&offsetedReader{r, 0}, binary.LittleEndian, &header); err != nil {
		return nil, errors.Wrap(err, "failed to read header")
	}
	var sp Sprite
	var err error
	switch header.Signature {
	default:
		return nil, errors.Errorf("bad signature; expected 0x9 or 0xf, got %#x", header.Signature)
	case 0x09:
		sp, err = newSprite8(r, header)
	case 0x0f:
		sp, err = newSprite32(r, header)
	case 0x19:
		sp, err = newSprite32Alpha(r, header)
	}
	return sp, err
}

type spriteHeader struct {
	Signature, FrameWidth, FrameHeight, FrameCount uint32
}
