package gosang

import "io"

type reader interface {
	Read([]byte) (int, error)
	ReadAt([]byte, int64) (int, error)
}

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
