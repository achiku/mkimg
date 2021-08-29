package main

import (
	"fmt"
	"image"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// DrawStringOpts options
type DrawStringOpts struct {
	ImageWidth       fixed.Int26_6
	ImageHeight      fixed.Int26_6
	VerticalMargin   fixed.Int26_6
	HorizontalMargin fixed.Int26_6
	FontSize         fixed.Int26_6
	LineSpace        fixed.Int26_6
	Verbose          bool
}

// DrawStringWrapped draw string wrapped
func DrawStringWrapped(d *font.Drawer, s string, opt *DrawStringOpts) {
	d.Dot.X = opt.HorizontalMargin
	d.Dot.Y = opt.FontSize + opt.VerticalMargin
	originalX, originalY := d.Dot.X, d.Dot.Y

	prevC := rune(-1)
	for _, c := range s {
		if prevC >= 0 {
			d.Dot.X += d.Face.Kern(prevC, c)
		}
		// dr, mask, maskp, advance, ok := d.Face.Glyph(d.Dot, c)
		advance, ok := d.Face.GlyphAdvance(c)
		if !ok {
			// TODO: is falling back on the U+FFFD glyph the responsibility of
			// the Drawer or the Face?
			// TODO: set prevC = '\ufffd'?
			continue
		}
		if d.Dot.X+advance >= (opt.ImageWidth - opt.HorizontalMargin*2) {
			if opt.Verbose {
				fmt.Printf("### new line: %#U, x=%d, y=%d, ", c, d.Dot.X, d.Dot.Y)
			}
			d.Dot.Y = originalY + d.Dot.Y + opt.LineSpace
			d.Dot.X = originalX
		}
		dr, mask, maskp, advance, ok := d.Face.Glyph(d.Dot, c)
		if opt.Verbose {
			fmt.Printf(
				"%#U: maskp=%+v, advance=%d, X=%d, w=%d, Y=%d, h=%d, realW=%d\n",
				c, maskp, advance, d.Dot.X, opt.ImageWidth, d.Dot.Y, opt.ImageHeight, (opt.ImageWidth - opt.HorizontalMargin*2))
		}
		draw.DrawMask(d.Dst, dr, d.Src, image.Point{}, mask, maskp, draw.Over)
		d.Dot.X += advance
		prevC = c
	}
}

// MeasureString returns how far dot would advance by drawing s with f.
func MeasureString(f font.Face, s string) (advance fixed.Int26_6) {
	prevC := rune(-1)
	for _, c := range s {
		if prevC >= 0 {
			advance += f.Kern(prevC, c)
		}
		a, ok := f.GlyphAdvance(c)
		if !ok {
			// TODO: is falling back on the U+FFFD glyph the responsibility of
			// the Drawer or the Face?
			// TODO: set prevC = '\ufffd'?
			continue
		}
		advance += a
		prevC = c
	}
	return advance
}
