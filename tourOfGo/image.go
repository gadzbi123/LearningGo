package main

import (
	"image"
	"image/color"
)

//	type IiiiImage interface{
//		ColorModel() color.Model
//		Bounds() image.Rectangle
//		At(x,y int) color.Color
//	}
type Image struct {
	values [][]struct{ r, g, b, a uint32 }
	h, w   int
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

type RGBAModel interface {
	Convert(c color.Color) color.Color
}

func (m Image) Convert(c color.Color) color.Color {
	c = m
	return c
}

func (i Image) RGBA() (uint32, uint32, uint32, uint32) {
	return 1, 2, 3, 4
}

func (i Image) ColorModel() RGBAModel {
	return i
}

func (i Image) At(x, y int) color.Color {
	return i
}

func main() {

	var m Image

	pic.ShowImage(m)
}
