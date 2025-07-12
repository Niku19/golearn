package main

import (
	"image"
	"image/color"
)

type Image struct {
	x, y int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.x, i.y)
}

func (i Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func exercise() {
	m := Image{x: 256, y: 256}
	// Implement a method to show m
	_ = m
}
