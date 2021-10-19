package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func TestDraw(t *testing.T) {
	m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	blue := color.RGBA{0, 0, 255, 255}

	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
	outfile, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	if err := png.Encode(outfile, m); err != nil {
		log.Fatal(err)
	}

}

func TestDrawStringWrapped(t *testing.T) {
	width, height := 1200, 630
	fontsize := 20.0
	txt := "あのイーハトーヴォのすきとおった風、夏でも底に冷たさをもつ青いそら、うつくしい森で飾られたモリーオ市、郊外のぎらぎらひかる草の波。"
	bk := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(bk, bk.Bounds(), image.White, image.ZP, draw.Src)

	ff := "./Koruri-Bold.ttf"
	fontBytes, err := ioutil.ReadFile(ff)
	if err != nil {
		t.Fatal(err)
	}
	ft, err := freetype.ParseFont(fontBytes)
	if err != nil {
		t.Fatal(err)
	}

	opt := truetype.Options{
		Size: fontsize,
	}
	face := truetype.NewFace(ft, &opt)
	dr := &font.Drawer{
		Dst:  bk,
		Src:  image.Black,
		Face: face,
		Dot:  fixed.Point26_6{},
	}
	t.Logf("dr.MeasureString(txt)=%d\n", MeasureString(dr.Face, txt))
	t.Logf("fixed.I(width)=%d\n", fixed.I(width))
	t.Logf("shift width=%d\n", int64(width<<6))

	dOpt := &DrawStringOpts{
		ImageWidth:       fixed.I(width),
		ImageHeight:      fixed.I(height),
		Verbose:          testing.Verbose(),
		FontSize:         fixed.I(int(fontsize)),
		LineSpace:        fixed.I(5),
		VerticalMargin:   fixed.I(10),
		HorizontalMargin: fixed.I(40),
	}

	DrawStringWrapped(dr, txt, dOpt)

	outfile, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	if err := png.Encode(outfile, bk); err != nil {
		log.Fatal(err)
	}
}

func TestCalculateInitialPoint(t *testing.T) {
	width, height := 1200, 630
	fontsize := 20.0
	txt := "あのイーハトーヴォのすきとおった風、夏でも底に冷たさをもつ青いそら、うつくしい森で飾られたモリーオ市、郊外のぎらぎらひかる草の波。"
	dOpt := &DrawStringOpts{
		ImageWidth:       fixed.I(width),
		ImageHeight:      fixed.I(height),
		Verbose:          testing.Verbose(),
		FontSize:         fixed.I(int(fontsize)),
		LineSpace:        fixed.I(5),
		VerticalMargin:   fixed.I(10),
		HorizontalMargin: fixed.I(40),
	}

	CalculateInitialPoint(txt, dOpt)
}
