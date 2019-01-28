package gosang

import (
	"image/color"
	"io"
)

// offsetedReader implements io.Reader combining io.ReaderAt and offset.
type offsetedReader struct {
	r      io.ReaderAt
	offset int64
}

func (or *offsetedReader) Read(p []byte) (int, error) {
	n, err := or.r.ReadAt(p, or.offset)
	or.offset += int64(n)
	return n, err
}

// sprite8Palette is a color palette used by 8-bit color sprites.
var sprite8Palette = color.Palette{
	color.RGBA{0x00, 0x00, 0x00, 0xff},
	color.RGBA{0x34, 0x5f, 0x2c, 0xff},
	color.RGBA{0x34, 0x51, 0x2c, 0xff},
	color.RGBA{0x34, 0x4a, 0x3f, 0xff},
	color.RGBA{0x2c, 0x3f, 0x37, 0xff},
	color.RGBA{0x2c, 0x42, 0x34, 0xff},
	color.RGBA{0x2c, 0x42, 0x37, 0xff},
	color.RGBA{0x29, 0x54, 0x25, 0xff},
	color.RGBA{0x25, 0x4a, 0x29, 0xff},
	color.RGBA{0x29, 0x3b, 0x29, 0xff},
	color.RGBA{0x25, 0x34, 0x30, 0xff},
	color.RGBA{0x25, 0x34, 0x37, 0xff},
	color.RGBA{0x21, 0x34, 0x34, 0xff},
	color.RGBA{0x21, 0x42, 0x1e, 0xff},
	color.RGBA{0x1e, 0x34, 0x1e, 0xff},
	color.RGBA{0x1a, 0x34, 0x29, 0xff},
	color.RGBA{0xd1, 0xaf, 0x74, 0xff},
	color.RGBA{0x88, 0x4d, 0x1e, 0xff},
	color.RGBA{0x77, 0x5f, 0x37, 0xff},
	color.RGBA{0xa8, 0x88, 0x54, 0xff},
	color.RGBA{0x99, 0x7b, 0x4a, 0xff},
	color.RGBA{0xbd, 0x9c, 0x5f, 0xff},
	color.RGBA{0x4a, 0x3b, 0x13, 0xff},
	color.RGBA{0x4d, 0x3f, 0x16, 0xff},
	color.RGBA{0x3b, 0x29, 0x0b, 0xff},
	color.RGBA{0x2c, 0x4d, 0x2c, 0xff},
	color.RGBA{0x42, 0x74, 0x37, 0xff},
	color.RGBA{0x42, 0x74, 0x34, 0xff},
	color.RGBA{0x42, 0x6a, 0x3b, 0xff},
	color.RGBA{0x3f, 0x66, 0x3b, 0xff},
	color.RGBA{0x3b, 0x5b, 0x34, 0xff},
	color.RGBA{0x37, 0x5f, 0x30, 0xff},
	color.RGBA{0xd1, 0x9c, 0x6d, 0xff},
	color.RGBA{0xb2, 0x77, 0x1a, 0xff},
	color.RGBA{0xd1, 0xbd, 0xb2, 0xff},
	color.RGBA{0xc3, 0x95, 0x21, 0xff},
	color.RGBA{0xb8, 0x7e, 0x16, 0xff},
	color.RGBA{0xc9, 0xac, 0x70, 0xff},
	color.RGBA{0xcc, 0xbd, 0xb5, 0xff},
	color.RGBA{0xc3, 0x9f, 0x51, 0xff},
	color.RGBA{0xb8, 0x74, 0x13, 0xff},
	color.RGBA{0xc0, 0x8f, 0x1a, 0xff},
	color.RGBA{0xc0, 0x99, 0x46, 0xff},
	color.RGBA{0xcc, 0xb5, 0x92, 0xff},
	color.RGBA{0xb2, 0x66, 0x0b, 0xff},
	color.RGBA{0xac, 0x51, 0x07, 0xff},
	color.RGBA{0x99, 0x30, 0x04, 0xff},
	color.RGBA{0x8f, 0x21, 0x04, 0xff},
	color.RGBA{0x77, 0x0b, 0x00, 0xff},
	color.RGBA{0x85, 0x7e, 0x70, 0xff},
	color.RGBA{0x88, 0x77, 0x63, 0xff},
	color.RGBA{0xb2, 0xa8, 0x9c, 0xff},
	color.RGBA{0x8c, 0x82, 0x70, 0xff},
	color.RGBA{0x9f, 0x95, 0x85, 0xff},
	color.RGBA{0xb2, 0xa8, 0x9c, 0xff},
	color.RGBA{0xbb, 0xb8, 0xaf, 0xff},
	color.RGBA{0x92, 0x8f, 0x82, 0xff},
	color.RGBA{0x99, 0x92, 0x88, 0xff},
	color.RGBA{0xbb, 0xaf, 0x9f, 0xff},
	color.RGBA{0x8c, 0x7b, 0x6a, 0xff},
	color.RGBA{0x8f, 0x82, 0x6d, 0xff},
	color.RGBA{0x37, 0x3b, 0x66, 0xff},
	color.RGBA{0x34, 0x37, 0x58, 0xff},
	color.RGBA{0x2c, 0x30, 0x51, 0xff},
	color.RGBA{0x92, 0x8c, 0x7b, 0xff},
	color.RGBA{0x3b, 0x30, 0x30, 0xff},
	color.RGBA{0x29, 0x16, 0x07, 0xff},
	color.RGBA{0x5f, 0x54, 0x4d, 0xff},
	color.RGBA{0x00, 0xaf, 0x00, 0xff},
	color.RGBA{0xaf, 0x00, 0x00, 0xff},
	color.RGBA{0xaf, 0xaf, 0x00, 0xff},
	color.RGBA{0x63, 0x58, 0x4d, 0xff},
	color.RGBA{0x3f, 0x37, 0x25, 0xff},
	color.RGBA{0x5b, 0x58, 0x4a, 0xff},
	color.RGBA{0x4d, 0x4a, 0x21, 0xff},
	color.RGBA{0x58, 0x46, 0x25, 0xff},
	color.RGBA{0x9c, 0x88, 0x6d, 0xff},
	color.RGBA{0x82, 0x74, 0x51, 0xff},
	color.RGBA{0x74, 0x63, 0x42, 0xff},
	color.RGBA{0x6d, 0x5f, 0x3f, 0xff},
	color.RGBA{0x8f, 0x7e, 0x5f, 0xff},
	color.RGBA{0x54, 0x4a, 0x37, 0xff},
	color.RGBA{0x6d, 0x63, 0x54, 0xff},
	color.RGBA{0x7b, 0x70, 0x5f, 0xff},
	color.RGBA{0x4d, 0x4a, 0x3b, 0xff},
	color.RGBA{0x77, 0x6d, 0x63, 0xff},
	color.RGBA{0x7e, 0x6a, 0x66, 0xff},
	color.RGBA{0x70, 0x6a, 0x5f, 0xff},
	color.RGBA{0x6a, 0x63, 0x58, 0xff},
	color.RGBA{0x3b, 0x37, 0x29, 0xff},
	color.RGBA{0x5b, 0x54, 0x42, 0xff},
	color.RGBA{0x4a, 0x46, 0x3b, 0xff},
	color.RGBA{0x6a, 0x63, 0x51, 0xff},
	color.RGBA{0x70, 0x6a, 0x5b, 0xff},
	color.RGBA{0x4d, 0x42, 0x30, 0xff},
	color.RGBA{0x42, 0x3b, 0x29, 0xff},
	color.RGBA{0x5f, 0x54, 0x42, 0xff},
	color.RGBA{0x42, 0x1e, 0x1a, 0xff},
	color.RGBA{0x46, 0x34, 0x13, 0xff},
	color.RGBA{0x54, 0x42, 0x1a, 0xff},
	color.RGBA{0x66, 0x42, 0x3b, 0xff},
	color.RGBA{0x51, 0x34, 0x2c, 0xff},
	color.RGBA{0x82, 0x58, 0x4d, 0xff},
	color.RGBA{0x74, 0x58, 0x30, 0xff},
	color.RGBA{0x85, 0x46, 0x1e, 0xff},
	color.RGBA{0x8f, 0x6d, 0x30, 0xff},
	color.RGBA{0x85, 0x5f, 0x29, 0xff},
	color.RGBA{0x99, 0x5b, 0x25, 0xff},
	color.RGBA{0x6a, 0x3b, 0x1a, 0xff},
	color.RGBA{0x66, 0x3b, 0x1a, 0xff},
	color.RGBA{0x70, 0x37, 0x1a, 0xff},
	color.RGBA{0x5f, 0x37, 0x16, 0xff},
	color.RGBA{0x77, 0x46, 0x1a, 0xff},
	color.RGBA{0x66, 0x13, 0x13, 0xff},
	color.RGBA{0x4d, 0x25, 0x1e, 0xff},
	color.RGBA{0xce, 0xb5, 0x99, 0xff},
	color.RGBA{0x5b, 0x30, 0x25, 0xff},
	color.RGBA{0x6a, 0x3b, 0x2c, 0xff},
	color.RGBA{0x70, 0x3f, 0x30, 0xff},
	color.RGBA{0x77, 0x42, 0x34, 0xff},
	color.RGBA{0x85, 0x4d, 0x51, 0xff},
	color.RGBA{0x8c, 0x51, 0x42, 0xff},
	color.RGBA{0x92, 0x58, 0x4d, 0xff},
	color.RGBA{0x99, 0x63, 0x4d, 0xff},
	color.RGBA{0xce, 0xc0, 0x9f, 0xff},
	color.RGBA{0xa8, 0x74, 0x5b, 0xff},
	color.RGBA{0xb5, 0x85, 0x6a, 0xff},
	color.RGBA{0xb8, 0x8c, 0x70, 0xff},
	color.RGBA{0x66, 0x8c, 0x8c, 0xff},
	color.RGBA{0x5f, 0x82, 0x7e, 0xff},
	color.RGBA{0x54, 0x70, 0x77, 0xff},
	color.RGBA{0x51, 0x77, 0x70, 0xff},
	color.RGBA{0x4d, 0x66, 0x6d, 0xff},
	color.RGBA{0x4a, 0x6a, 0x70, 0xff},
	color.RGBA{0x42, 0x66, 0x66, 0xff},
	color.RGBA{0x4a, 0x6a, 0x66, 0xff},
	color.RGBA{0x4a, 0x58, 0x7e, 0xff},
	color.RGBA{0x3f, 0x4a, 0x70, 0xff},
	color.RGBA{0x3f, 0x63, 0x5f, 0xff},
	color.RGBA{0x3f, 0x58, 0x5f, 0xff},
	color.RGBA{0x37, 0x42, 0x66, 0xff},
	color.RGBA{0x34, 0x4a, 0x5b, 0xff},
	color.RGBA{0x30, 0x37, 0x5b, 0xff},
	color.RGBA{0x25, 0x2c, 0x46, 0xff},
	color.RGBA{0xb5, 0xc0, 0xcc, 0xff},
	color.RGBA{0x34, 0x42, 0x5f, 0xff},
	color.RGBA{0x30, 0x3f, 0x5b, 0xff},
	color.RGBA{0x77, 0x5f, 0x51, 0xff},
	color.RGBA{0x58, 0x4a, 0x3b, 0xff},
	color.RGBA{0x51, 0x42, 0x34, 0xff},
	color.RGBA{0x58, 0x46, 0x37, 0xff},
	color.RGBA{0xaf, 0x99, 0x4d, 0xff},
	color.RGBA{0xac, 0x95, 0x4a, 0xff},
	color.RGBA{0x63, 0x54, 0x13, 0xff},
	color.RGBA{0x7b, 0x6a, 0x13, 0xff},
	color.RGBA{0x85, 0x70, 0x13, 0xff},
	color.RGBA{0x54, 0x5f, 0x5f, 0xff},
	color.RGBA{0x3f, 0x4a, 0x4a, 0xff},
	color.RGBA{0x85, 0xb5, 0xaf, 0xff},
	color.RGBA{0x70, 0x9c, 0x9c, 0xff},
	color.RGBA{0x4a, 0x51, 0x58, 0xff},
	color.RGBA{0x7b, 0x29, 0x1a, 0xff},
	color.RGBA{0x88, 0x34, 0x21, 0xff},
	color.RGBA{0x9c, 0x42, 0x29, 0xff},
	color.RGBA{0xac, 0x4d, 0x34, 0xff},
	color.RGBA{0xbb, 0x63, 0x3f, 0xff},
	color.RGBA{0xc9, 0x85, 0x58, 0xff},
	color.RGBA{0x82, 0x2c, 0x1e, 0xff},
	color.RGBA{0x92, 0x3b, 0x25, 0xff},
	color.RGBA{0xa2, 0x46, 0x2c, 0xff},
	color.RGBA{0xb5, 0x5b, 0x3b, 0xff},
	color.RGBA{0xc3, 0x70, 0x4a, 0xff},
	color.RGBA{0xd1, 0xa5, 0x6d, 0xff},
	color.RGBA{0xd4, 0xd4, 0xb8, 0xff},
	color.RGBA{0xd4, 0xc9, 0x9f, 0xff},
	color.RGBA{0xd4, 0xc0, 0x92, 0xff},
	color.RGBA{0xd4, 0xb5, 0x88, 0xff},
	color.RGBA{0xd4, 0xaf, 0x82, 0xff},
	color.RGBA{0xd4, 0xc6, 0x95, 0xff},
	color.RGBA{0xc9, 0x7b, 0x51, 0xff},
	color.RGBA{0xce, 0x8c, 0x5f, 0xff},
	color.RGBA{0xd1, 0xa5, 0x74, 0xff},
	color.RGBA{0xd1, 0xaf, 0x74, 0xff},
	color.RGBA{0xd1, 0x9c, 0x6d, 0xff},
	color.RGBA{0x34, 0x51, 0x37, 0xff},
	color.RGBA{0x51, 0x25, 0x04, 0xff},
	color.RGBA{0x4a, 0x21, 0x04, 0xff},
	color.RGBA{0x42, 0x1e, 0x00, 0xff},
	color.RGBA{0x37, 0x1a, 0x00, 0xff},
	color.RGBA{0x34, 0x16, 0x00, 0xff},
	color.RGBA{0x2c, 0x13, 0x00, 0xff},
	color.RGBA{0x1e, 0x0b, 0x00, 0xff},
	color.RGBA{0x63, 0x2c, 0x04, 0xff},
	color.RGBA{0x51, 0x66, 0x85, 0xff},
	color.RGBA{0x4a, 0x58, 0x7b, 0xff},
	color.RGBA{0x46, 0x6d, 0x66, 0xff},
	color.RGBA{0x42, 0x51, 0x74, 0xff},
	color.RGBA{0x3b, 0x46, 0x66, 0xff},
	color.RGBA{0x37, 0x4d, 0x5b, 0xff},
	color.RGBA{0x2c, 0x37, 0x54, 0xff},
	color.RGBA{0x25, 0x2c, 0x4d, 0xff},
	color.RGBA{0x1e, 0x21, 0x46, 0xff},
	color.RGBA{0x1a, 0x1e, 0x42, 0xff},
	color.RGBA{0x1a, 0x1e, 0x42, 0xff},
	color.RGBA{0x16, 0x1a, 0x3b, 0xff},
	color.RGBA{0x16, 0x16, 0x37, 0xff},
	color.RGBA{0x0f, 0x0f, 0x34, 0xff},
	color.RGBA{0x58, 0x29, 0x04, 0xff},
	color.RGBA{0x25, 0x25, 0x25, 0xff},
	color.RGBA{0x46, 0x46, 0x46, 0xff},
	color.RGBA{0x6d, 0x6d, 0x6d, 0xff},
	color.RGBA{0x8f, 0x8f, 0x8f, 0xff},
	color.RGBA{0x25, 0x13, 0x00, 0xff},
	color.RGBA{0x46, 0x2c, 0x00, 0xff},
	color.RGBA{0x6d, 0x4d, 0x00, 0xff},
	color.RGBA{0x8f, 0x5f, 0x00, 0xff},
	color.RGBA{0x25, 0x00, 0x25, 0xff},
	color.RGBA{0x3f, 0x00, 0x3f, 0xff},
	color.RGBA{0x5b, 0x00, 0x5b, 0xff},
	color.RGBA{0x77, 0x00, 0x77, 0xff},
	color.RGBA{0x00, 0x25, 0x25, 0xff},
	color.RGBA{0x00, 0x3f, 0x3f, 0xff},
	color.RGBA{0x00, 0x5b, 0x5b, 0xff},
	color.RGBA{0x00, 0x77, 0x77, 0xff},
	color.RGBA{0x25, 0x25, 0x00, 0xff},
	color.RGBA{0x46, 0x46, 0x00, 0xff},
	color.RGBA{0x6d, 0x6d, 0x00, 0xff},
	color.RGBA{0x8f, 0x8f, 0x00, 0xff},
	color.RGBA{0x00, 0x25, 0x00, 0xff},
	color.RGBA{0x00, 0x3f, 0x00, 0xff},
	color.RGBA{0x00, 0x5b, 0x00, 0xff},
	color.RGBA{0x00, 0x77, 0x00, 0xff},
	color.RGBA{0x25, 0x00, 0x00, 0xff},
	color.RGBA{0x4a, 0x00, 0x00, 0xff},
	color.RGBA{0x6d, 0x00, 0x00, 0xff},
	color.RGBA{0x8f, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0x25, 0xff},
	color.RGBA{0x00, 0x00, 0x4a, 0xff},
	color.RGBA{0x00, 0x00, 0x6d, 0xff},
	color.RGBA{0x00, 0x00, 0x8f, 0xff},
	color.RGBA{0xc9, 0xc9, 0xc9, 0xff},
	color.RGBA{0x0f, 0x0f, 0x0f, 0xff},
	color.RGBA{0x1e, 0x1e, 0x1e, 0xff},
	color.RGBA{0x2c, 0x2c, 0x2c, 0xff},
	color.RGBA{0x3f, 0x3f, 0x3f, 0xff},
	color.RGBA{0x4d, 0x4d, 0x4d, 0xff},
	color.RGBA{0x5b, 0x5b, 0x5b, 0xff},
	color.RGBA{0x6a, 0x6a, 0x6a, 0xff},
	color.RGBA{0x7b, 0x7b, 0x7b, 0xff},
	color.RGBA{0x88, 0x88, 0x88, 0xff},
	color.RGBA{0x95, 0x95, 0x95, 0xff},
	color.RGBA{0xa2, 0xa2, 0xa2, 0xff},
	color.RGBA{0xb2, 0xb2, 0xb2, 0xff},
	color.RGBA{0xbd, 0xbd, 0xbd, 0xff},
	color.RGBA{0xd4, 0xc0, 0xd4, 0xff},
	color.RGBA{0xd4, 0xd4, 0xd4, 0xff},
}
