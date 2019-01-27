package gosang

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
)

// Sprite represents single sprite.
type Sprite interface {
	Width() int
	Height() int
	Count() int
}

// NewSprite creates new sprite from r.
func NewSprite(r io.ReadSeeker) (Sprite, error) {
	var header struct {
		Signature uint32
		Width     uint32
		Height    uint32
		Count     uint32
	}
	if err := binary.Read(r, binary.LittleEndian, &header); err != nil {
		return nil, errors.Wrap(err, "failed to read header")
	}
	var sp Sprite
	switch header.Signature {
	default:
		return nil, errors.Errorf("bad signature; expected 0x9 or 0xf, got %#x", header.Signature)
	case 0x9:
		sp = &Sprite8{
			width:  int(header.Width),
			height: int(header.Height),
			count:  int(header.Count),
		}
	case 0xf:
		sp = &Sprite32{
			width:  int(header.Width),
			height: int(header.Height),
			count:  int(header.Count),
		}
	}
	return sp, nil
}
