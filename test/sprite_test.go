package test

import (
	"os"
	"testing"

	"github.com/hallazzang/gosang"
)

func TestNewSprite(t *testing.T) {
	for _, tc := range []struct {
		path                            string
		colorBits, width, height, count int
	}{
		{`data\arrow.spr`, 8, 20, 20, 10},
		{`data\BUTTMENU_ONLINE_1.S32`, 32, 24, 52, 2},
	} {
		func() {
			f, err := os.Open(tc.path)
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
			if w := sp.Width(); w != tc.width {
				t.Errorf("bad sprite frame width; expected %d, got %d", tc.width, w)
			}
			if h := sp.Height(); h != tc.height {
				t.Errorf("bad sprite frame height; expected %d, got %d", tc.height, h)
			}
			if c := sp.Count(); c != tc.count {
				t.Errorf("bad sprite frame count; expected %d, got %d", tc.count, c)
			}
		}()
	}
}
