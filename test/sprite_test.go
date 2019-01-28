package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/hallazzang/gosang"
)

func TestNewSprite(t *testing.T) {
	for _, tc := range []struct {
		path                                           string
		colorBits, frameWidth, frameHeight, frameCount int
	}{
		{"arrow.spr", 8, 20, 20, 10},
		{"BUTTMENU_ONLINE_1.S32", 32, 24, 52, 2},
	} {
		func() {
			f, err := os.Open(filepath.Join("data", tc.path))
			if err != nil {
				t.Fatalf("failed to open sprite file: %v", err)
			}
			defer f.Close()
			sp, err := gosang.NewSprite(f)
			if err != nil {
				t.Fatalf("failed to open sprite: %v", err)
			}
			if cb := sp.ColorBits(); cb != tc.colorBits {
				t.Errorf("bad sprite color bits; expected %d, got %d", tc.colorBits, cb)
			}
			if w := sp.FrameWidth(); w != tc.frameWidth {
				t.Errorf("bad sprite frame width; expected %d, got %d", tc.frameWidth, w)
			}
			if h := sp.FrameHeight(); h != tc.frameHeight {
				t.Errorf("bad sprite frame height; expected %d, got %d", tc.frameHeight, h)
			}
			if c := sp.FrameCount(); c != tc.frameCount {
				t.Errorf("bad sprite frame count; expected %d, got %d", tc.frameCount, c)
			}
		}()
	}
}
