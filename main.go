package main

import (
	"flag"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	txt = flag.String("txt", "hello, world!", "text to image")
	out = flag.String("outfile", "out.png", "output image path")
	fs  = flag.Float64("fontsize", 100.0, "font size")
	ff  = flag.String("fontfile", "", "ttf font file path")
	w   = flag.Int("width", 1200, "image width")
	h   = flag.Int("height", 630, "image height")
)

func main() {
	flag.Parse()

	fontsize := *fs
	width, height := *w, *h

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), image.White, image.ZP, draw.Src)
	fontBytes, err := ioutil.ReadFile(*ff)
	if err != nil {
		log.Fatal(err)
	}
	ft, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Fatal(err)
	}

	opt := truetype.Options{
		Size: fontsize,
	}
	face := truetype.NewFace(ft, &opt)
	dr := &font.Drawer{
		Dst:  img,
		Src:  image.Black,
		Face: face,
		Dot:  fixed.Point26_6{},
	}
	x := (fixed.I(width) - dr.MeasureString(*txt)) / 2
	dr.Dot.X = x
	y := (height + int(fontsize)/2) / 2
	dr.Dot.Y = fixed.I(y)

	dr.DrawString(*txt)

	outfile, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	if err := png.Encode(outfile, img); err != nil {
		log.Fatal(err)
	}
}
