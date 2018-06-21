package main

import (
	"image"
	"tour/pic"
	"image/color"
)

type Image struct{
	Width, Height int
	c uint8
}

func (r *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.Width, r.Height)
}

func (r *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (r *Image) At(x, y int) color.Color {
	return color.RGBA{r.c+uint8(x), r.c+uint8(y), 255, 255}
}

func main() {
	m := Image{50, 50, 128}
	pic.ShowImage(&m)
}
