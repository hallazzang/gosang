package gosang

import (
	"bufio"
	"encoding/binary"
	"image"
	"image/color"
	"io"

	"github.com/pkg/errors"
)

// sprite32 is a 32-bit color sprite.
type sprite32 struct {
	r       io.ReaderAt
	width   int
	height  int
	count   int
	offsets []uint32
}

func newSprite32(r io.ReaderAt, header spriteHeader) (*sprite32, error) {
	sp := &sprite32{
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

func (sp *sprite32) ColorBits() int {
	return 32
}

func (sp *sprite32) Width() int {
	return sp.width
}

func (sp *sprite32) Height() int {
	return sp.height
}

func (sp *sprite32) Count() int {
	return sp.count
}

func (sp *sprite32) Frame(idx int) (image.Image, error) {
	if idx < 0 || idx > sp.count-1 {
		return nil, errors.New("frame index out of range")
	}
	img := image.NewRGBA(image.Rect(0, 0, sp.width, sp.height))
	r := bufio.NewReader(&offsetedReader{sp.r, 0xe4c + int64(sp.offsets[idx])})
	for y := 0; y < sp.height; y++ {
		for x := 0; x < sp.width; {
			var p struct{ Count, Blue, Green, Red byte }
			if err := binary.Read(r, binary.LittleEndian, &p); err != nil {
				return nil, errors.Wrap(err, "failed to read frame data")
			}
			for ; p.Count > 0; p.Count-- {
				img.SetRGBA(x, y, color.RGBA{p.Red, p.Green, p.Blue, 0xff})
				x++
			}
		}
	}
	return img, nil
}
