package main

import (
	"fmt"
	"image"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func drawStringDebug(s string, d *font.Drawer) {
	prevC := rune(-1)
	for _, c := range s {
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
		fmt.Printf("%#U: maskp=%v, advance=%d, X=%d, Y=%d\n", c, maskp, advance, d.Dot.X, d.Dot.Y)
		prevC = c
	}
}

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
	// originalX := d.Dot.X
	originalY := d.Dot.Y
	prevC := rune(-1)
	for _, c := range s {
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
		if d.Dot.X+advance >= (opt.ImageWidth - opt.HorizontalMargin*2) {
			d.Dot.Y = originalY + d.Dot.Y + opt.LineSpace
			d.Dot.X = fixed.I(0)
		}

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

// Text represents text input
type Text struct {
	OriginalText string
	Lines        []TextLine
	Padding      int
	Width        int
}

// TextLine text line
type TextLine struct {
	StartingX int
	StartingY int
	Words     string
}

func calculatePoints(s string, w, h int) (int, int) {
	return 0, 0
}
