package gosang

type reader interface {
	Read([]byte) (int, error)
	ReadAt([]byte, int64) (int, error)
}
