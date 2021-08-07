package main

import (
	"fmt"
	"image"
	"image/draw"

	"golang.org/x/image/font"
)

func drawStringDebug(s string, d *font.Drawer) {
	prevC := rune(-1)
	for _, c := range s {
		fmt.Printf("%#U\n", c)
		if prevC >= 0 {
			d.Dot.X += d.Face.Kern(prevC, c)
		}
		dr, mask, maskp, advance, ok := d.Face.Glyph(d.Dot, c)
		if !ok {
			// TODO: is falling back on the U+FFFD glyph the responsibility of
			// the Drawer or the Face?
			// TODO: set prevC = '\ufffd'?
			continue
		}
		draw.DrawMask(d.Dst, dr, d.Src, image.Point{}, mask, maskp, draw.Over)
		d.Dot.X += advance
		fmt.Printf("maskp=%v, advance=%d, X=%d\n", maskp, advance, d.Dot.X)
		prevC = c
	}
}

func calculatePoints(s string, w, h int) (int, int) {
	return 0, 0
}
