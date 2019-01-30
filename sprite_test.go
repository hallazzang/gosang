package gosang

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewSprite(t *testing.T) {
	for _, tc := range []struct {
		path        string
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
			f, err := os.Open(filepath.Join("test", "data", tc.path))
			if err != nil {
				t.Fatalf("failed to open sprite file: %v", err)
			}
			defer f.Close()
			sp, err := NewSprite(f)
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
			if w := sp.Width(); w != tc.width {
				t.Errorf("bad sprite width; expected %d, got %d", tc.width, w)
			}
			if h := sp.Height(); h != tc.height {
				t.Errorf("bad sprite height; expected %d, got %d", tc.height, h)
			}
		}()
	}
}

func TestFrameSizeAndOffset(t *testing.T) {
	for _, tc := range []string{"arrow.spr", "BUTTMENU_ONLINE_1.S32", "WindCutter.S32"} {
		func() {
			f, err := os.Open(filepath.Join("test", "data", tc))
			if err != nil {
				t.Fatalf("failed to open sprite file: %v", err)
			}
			defer f.Close()
			sp, err := NewSprite(f)
			if err != nil {
				t.Fatal(err)
			}
			ao := int64(0)
			for i := 0; i < sp.FrameCount(); i++ {
				o, err := sp.frameOffset(i)
				if err != nil {
					t.Errorf("faild to get frame #%d's offset: %v", i, err)
				}
				if o != ao {
					t.Errorf("frame #%d's offset is incorrect; expected %d, got %d", i, ao, o)
				}
				s, err := sp.frameSize(i)
				if err != nil {
					t.Errorf("failed to get frame #%d's size: %v", i, err)
				}
				ao += int64(s)
			}
		}()
	}
}
