package models

import (
	"image/color"
)

type BarPos int

const (
	BarTop BarPos = iota
	BarBottom
)

type Bar struct {
	Position   BarPos
	Width      int
	Height     int
	Background color.Color
	TextColor  color.Color
}

func NewBar(width, height int, background color.Color) *Bar {
	return &Bar{
		Position:   BarTop,
		Width:      width,
		Height:     height,
		Background: background,
		TextColor:  nil,
	}
}
