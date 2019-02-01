package gosang

import (
	"encoding/binary"
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
	Frame(idx int) (*Frame, error) // Specific frame's data.
	Save(w io.Writer) error        // Write sprite data to w.

	frameOffset(idx int) (int64, error)
	frameSize(idx int) (int, error)
	loadFrame(idx int) (*Frame, error)
}

// OpenSprite creates new sprite from r. It can accept all three type of
// sprites: 8-bit sprite(.spr), 32-bit sprite w/o alpha channel, 32-bit
// sprite w/ alpha channel.
func OpenSprite(r io.ReaderAt) (Sprite, error) {
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
	for i := uint32(0); i < header.FrameCount; i++ {
		if _, err := sp.loadFrame(int(i)); err != nil {
			return nil, errors.Wrapf(err, "failed to load frame #%d", i)
		}
	}
	return sp, err
}

type sprite struct {
	r           io.ReaderAt
	frameWidth  uint32
	frameHeight uint32
	frameCount  uint32
	offsets     []uint32
	width       uint32
	height      uint32
	lastOffset  uint32
	frames      []*Frame
}

func (sp *sprite) FrameWidth() int {
	return int(sp.frameWidth)
}

func (sp *sprite) FrameHeight() int {
	return int(sp.frameHeight)
}

func (sp *sprite) FrameCount() int {
	return int(sp.frameCount)
}

func (sp *sprite) Width() int {
	return int(sp.width)
}

func (sp *sprite) Height() int {
	return int(sp.height)
}

func (sp *sprite) Frame(idx int) (*Frame, error) {
	if idx < 0 || idx > int(sp.frameCount-1) {
		return nil, errors.New("frame index out of range")
	}
	return sp.frames[idx], nil
}

func (sp *sprite) frameOffset(idx int) (int64, error) {
	if idx < 0 || idx > int(sp.frameCount-1) {
		return 0, errors.New("frame index out of range")
	}
	return int64(sp.offsets[idx]), nil
}

func (sp *sprite) frameSize(idx int) (int, error) {
	if idx < 0 || idx > int(sp.frameCount-1) {
		return 0, errors.New("frame index out of range")
	} else if idx < int(sp.frameCount-1) {
		return int(sp.offsets[idx+1] - sp.offsets[idx]), nil
	}
	if sp.lastOffset == 0 {
		if err := binary.Read(&offsetedReader{sp.r, 0xe20}, binary.LittleEndian, &sp.lastOffset); err != nil {
			return 0, errors.Wrap(err, "failed to read sprite's last data offset")
		}
	}
	return int(sp.lastOffset - sp.offsets[idx]), nil
}

type spriteHeader struct {
	Signature, FrameWidth, FrameHeight, FrameCount uint32
}
