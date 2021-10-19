package main

import (
	"flag"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	txt   = flag.String("txt", "hello, world!", "text to image")
	out   = flag.String("outfile", "out.png", "output image path")
	bg    = flag.String("background", "", "background image path")
	fs    = flag.Float64("fontsize", 100.0, "font size")
	ff    = flag.String("fontfile", "", "ttf font file path")
	w     = flag.Int("width", 1200, "image width")
	h     = flag.Int("height", 630, "image height")
	space = flag.Bool("space", false, "space image")
	debug = flag.Bool("debug", false, "debug output")
)

func main() {
	flag.Parse()

	fontsize := *fs
	width, height := *w, *h

	var (
		bk    draw.Image
		img   draw.Image
		ok    bool
		color *image.Uniform
	)
	if *space {
		sf, err := os.Open(path.Join("templates", "space2.png"))
		if err != nil {
			log.Fatal(err)
		}
		dimg, _, err := image.Decode(sf)
		if err != nil {
			log.Fatalf("image.Decode failed: %s", err)
		}

		img, ok = dimg.(draw.Image)
		if !ok {
			log.Fatal(err)
		}
		bk = image.NewRGBA(image.Rect(0, 0, width, height))
		draw.Draw(bk, img.Bounds(), img, image.ZP, draw.Src)
		color = image.White
	} else if *bg != "" {
		sf, err := os.Open(*bg)
		if err != nil {
			log.Fatal(err)
		}
		dimg, _, err := image.Decode(sf)
		if err != nil {
			log.Fatalf("image.Decode failed: %s", err)
		}

		img, ok = dimg.(draw.Image)
		if !ok {
			log.Fatal(err)
		}
		bk = image.NewRGBA(image.Rect(0, 0, width, height))
		draw.Draw(bk, img.Bounds(), img, image.ZP, draw.Src)
		color = image.White
	} else {
		bk = image.NewRGBA(image.Rect(0, 0, width, height))
		draw.Draw(bk, bk.Bounds(), image.White, image.ZP, draw.Src)
		color = image.Black
	}

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
		Dst:  bk,
		Src:  color,
		Face: face,
		Dot:  fixed.Point26_6{},
	}
	// x := (fixed.I(width) - dr.MeasureString(*txt)) / 2
	// dr.Dot.X = x
	// y := (height + int(fontsize)/2) / 2
	// dr.Dot.Y = fixed.I(y)

	dOpt := &DrawStringOpts{
		ImageWidth:       fixed.I(width),
		ImageHeight:      fixed.I(height),
		Verbose:          *debug,
		FontSize:         fixed.I(int(fontsize)),
		LineSpace:        fixed.I(5),
		VerticalMargin:   fixed.I(10),
		HorizontalMargin: fixed.I(40),
	}

	DrawStringWrapped(dr, *txt, dOpt)

	outfile, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	if err := png.Encode(outfile, bk); err != nil {
		log.Fatal(err)
	}
}
