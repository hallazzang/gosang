package gosang

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewSprite(t *testing.T) {
	for _, tc := range []struct {
		name        string
		colorBits   int
		frameWidth  int
		frameHeight int
		frameCount  int
		width       int
		height      int
	}{
		{"arrow.spr", 8, 20, 20, 10, 200, 20},
		{"BUTTMENU_ONLINE_1.S32", 32, 24, 52, 2, 48, 52},
		{"WindCutter.S32", 32, 640, 480, 15, 9600, 480},
	} {
		func() {
			f, err := os.Open(filepath.Join("test", "data", tc.name))
			if err != nil {
				t.Fatalf("sprite %q: failed to open file: %v", tc.name, err)
			}
			defer f.Close()
			sp, err := NewSprite(f)
			if err != nil {
				t.Fatalf("sprite %q: failed to open sprite: %v", tc.name, err)
			}
			if cb := sp.ColorBits(); cb != tc.colorBits {
				t.Errorf("sprite %q: bad color bits; expected %d, got %d", tc.name, tc.colorBits, cb)
			}
			if w := sp.FrameWidth(); w != tc.frameWidth {
				t.Errorf("sprite %q: bad frame width; expected %d, got %d", tc.name, tc.frameWidth, w)
			}
			if h := sp.FrameHeight(); h != tc.frameHeight {
				t.Errorf("sprite %q: bad frame height; expected %d, got %d", tc.name, tc.frameHeight, h)
			}
			if c := sp.FrameCount(); c != tc.frameCount {
				t.Errorf("sprite %q: bad frame count; expected %d, got %d", tc.name, tc.frameCount, c)
			}
			if w := sp.Width(); w != tc.width {
				t.Errorf("sprite %q: bad width; expected %d, got %d", tc.name, tc.width, w)
			}
			if h := sp.Height(); h != tc.height {
				t.Errorf("sprite %q: bad height; expected %d, got %d", tc.name, tc.height, h)
			}
		}()
	}
}

func TestFrameSizeAndOffset(t *testing.T) {
	for _, name := range []string{"arrow.spr", "BUTTMENU_ONLINE_1.S32", "WindCutter.S32"} {
		func() {
			f, err := os.Open(filepath.Join("test", "data", name))
			if err != nil {
				t.Fatalf("sprite %q: failed to open file: %v", name, err)
			}
			defer f.Close()
			sp, err := NewSprite(f)
			if err != nil {
				t.Fatalf("sprite %q: failed to open sprite: %v", name, err)
			}
			ao := int64(0)
			for i := 0; i < sp.FrameCount(); i++ {
				o, err := sp.frameOffset(i)
				if err != nil {
					t.Errorf("sprite %q: faild to get frame #%d's offset: %v", name, i, err)
				}
				if o != ao {
					t.Errorf("sprite %q: frame #%d's offset is incorrect; expected %d, got %d", name, i, ao, o)
				}
				s, err := sp.frameSize(i)
				if err != nil {
					t.Errorf("sprite %q: failed to get frame #%d's size: %v", name, i, err)
				}
				ao += int64(s)
			}
		}()
	}
}
