package test

import (
	"os"
	"testing"

	"github.com/hallazzang/gosang"
)

func TestNewSprite(t *testing.T) {
	f, err := os.Open(`data\arrow.spr`)
	if err != nil {
		t.Fatalf("failed to open sprite file: %v", err)
	}
	sp, err := gosang.NewSprite(f)
	if err != nil {
		t.Fatalf("failed to create new sprite: %v", err)
	}
	if w := sp.Width(); w != 20 {
		t.Errorf("bad sprite frame width; expected 20, got %d", w)
	}
	if h := sp.Height(); h != 20 {
		t.Errorf("bad sprite frame height; expected 20, got %d", h)
	}
	if c := sp.Count(); c != 10 {
		t.Errorf("bad sprite frame count; expected 10, got %d", c)
	}
}
