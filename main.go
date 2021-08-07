package main

import (
	"flag"

	"github.com/fogleman/gg"
)

var (
	txt = flag.String("txt", "hello, world!", "text to image")
	out = flag.String("outfile", "out.png", "output image path")
	fs  = flag.Float64("fontsize", 100.0, "font size")
	ff  = flag.String("fontfile", "", "ttf font file path")
	w   = flag.Int("width", 1200, "image width")
	h   = flag.Int("height", 630, "image height")
	// space = flag.Bool("space", false, "space image")
)

func main() {
	flag.Parse()

	W := 1200.0
	H := 630.0
	P := 16.0
	dc := gg.NewContext(int(W), int(H))
	dc.SetRGB(1, 1, 1)
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace(*ff, 18); err != nil {
		panic(err)
	}
	dc.DrawStringWrapped(*txt, W/2-P, H/2-P, 1, 1, W, 1.75, gg.AlignCenter)
	dc.SavePNG(*out)
}
