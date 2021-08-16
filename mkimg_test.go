package main

import (
	"fmt"
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

func TestDrawString(t *testing.T) {
	width, height := 1200, 630
	fontsize := 10.0
	txt := "あのイーハトーヴォのすきとおった風、夏でも底に冷たさをもつ青いそら、うつくしい森で飾られたモリーオ市、郊外のぎらぎらひかる草の波。"
	// txt := "それは本当にそう"
	// txt := "Call me Ishmael. Some years ago—never mind how long precisely—having little or no money in my purse, and nothing particular to interest me on shore, I thought I would sail about a little and see the watery part of the world. It is a way I have of driving off the spleen and regulating the circulation. Whenever I find myself growing grim about the mouth; whenever it is a damp, drizzly November in my soul; whenever I find myself involuntarily pausing before coffin warehouses, and bringing up the rear of every funeral I meet; and especially whenever my hypos get such an upper hand of me, that it requires a strong moral principle to prevent me from deliberately stepping into the street, and methodically knocking people's hats off—then, I account it high time to get to sea as soon as I can. This is my substitute for pistol and ball. With a philosophical flourish Cato throws himself upon his sword; I quietly take to the ship. There is nothing surprising in this. If they but knew it, almost all men in their degree, some time or other, cherish very nearly the same feelings towards the ocean with me."
	ff := "./Koruri-Bold.ttf"
	bk := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(bk, bk.Bounds(), image.White, image.ZP, draw.Src)
	color := image.Black

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
		Src:  color,
		Face: face,
		Dot:  fixed.Point26_6{},
	}
	fmt.Printf("dr.MeasureString(txt)=%d\n", MeasureString(dr.Face, txt))
	fmt.Printf("fixed.I(width)=%d\n", fixed.I(width))
	fmt.Printf("shift width=%d\n", int64(width<<6))
	//x := (fixed.I(width) - dr.MeasureString(txt)) / 2
	//dr.Dot.X = x
	y := (height + int(fontsize)/2) / 2
	dr.Dot.Y = fixed.I(y)

	drawStringDebug(txt, dr)

	outfile, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	if err := png.Encode(outfile, bk); err != nil {
		log.Fatal(err)
	}
}

func TestDrawStringWrapped(t *testing.T) {
	width, height := 1200, 630
	fontsize := 50.0
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
	fmt.Printf("dr.MeasureString(txt)=%d\n", MeasureString(dr.Face, txt))
	fmt.Printf("fixed.I(width)=%d\n", fixed.I(width))
	fmt.Printf("shift width=%d\n", int64(width<<6))

	dOpt := &DrawStringOpts{
		ImageWidth:       fixed.I(width),
		ImageHeight:      fixed.I(height),
		Verbose:          true,
		FontSize:         fixed.I(int(fontsize)),
		LineSpace:        fixed.I(5),
		VerticalMargin:   fixed.I(10),
		HorizontalMargin: fixed.I(80),
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
