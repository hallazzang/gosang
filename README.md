# gosang

[![godoc]](https://godoc.org/github.com/hallazzang/gosang)
[![goreportcard]](https://goreportcard.com/report/github.com/hallazzang/gosang)

Gersang sprite library for Go

## Example

Here's a simple example that uses **gosang** to extract frame images from sprite
to files in PNG format.

It works well for both 8-bit sprite(.spr) and 32-bit(.S32) sprite.

```go
package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/hallazzang/gosang"
)

func main() {
	f, err := os.Open(`data\BUTTMENU_ONLINE_1.S32`)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	sp, err := gosang.OpenSprite(f)
	if err != nil {
		panic(err)
	}
	for i := 0; i < sp.FrameCount(); i++ {
		frame, err := sp.Frame(i)
		if err != nil {
			panic(err)
		}
		func() {
			out, err := os.Create(fmt.Sprintf("output%d.png", i))
			if err != nil {
				panic(err)
			}
			defer out.Close()
			if err := png.Encode(out, frame.Image()); err != nil {
				panic(err)
			}
		}()
	}
}
```

[godoc]: https://godoc.org/github.com/hallazzang/gosang?status.svg
[goreportcard]: https://goreportcard.com/badge/github.com/hallazzang/gosang
