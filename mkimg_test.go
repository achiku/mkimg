package main

import (
	"fmt"
	"image"
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

func TestDrawString(t *testing.T) {
	width, height := 1200, 630
	fontsize := 10.0
	txt := "あのイーハトーヴォのすきとおった風、夏でも底に冷たさをもつ青いそら、うつくしい森で飾られたモリーオ市、郊外のぎらぎらひかる草の波。"
	// txt := "それは本当にそう"
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
	fmt.Printf("dr.MeasureString(txt)=%d\n", dr.MeasureString(txt))
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
