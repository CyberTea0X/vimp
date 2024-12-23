package models

import (
	"image/color"
	"vimp/helpers"
)

type CurrentColor = int

const (
	ColorPrimary CurrentColor = iota
	ColorSecondary
)

type Palette struct {
	BackgroundColor color.Color
	PrimaryColor    color.Color
	SecondaryColor  color.Color
	CurrentColor    CurrentColor
	Color3          color.Color
	Color4          color.Color
	Color5          color.Color
	Color6          color.Color
	Color7          color.Color
	Color8          color.Color
	Color9          color.Color
}

func (p *Palette) GetCurrentColor() color.Color {
	switch p.CurrentColor {
	case ColorPrimary:
		return p.PrimaryColor
	case ColorSecondary:
		return p.SecondaryColor
	}
	return color.RGBA{}
}

func (p *Palette) SetCurrentColor(c color.Color) {
	switch p.CurrentColor {
	case ColorPrimary:
		p.PrimaryColor = c
	case ColorSecondary:
		p.SecondaryColor = c
	}
}

func DefaultPalette() *Palette {
	return &Palette{
		BackgroundColor: color.White,
		PrimaryColor:    color.Black,
		SecondaryColor:  color.White,
		Color3:          helpers.RGBFromInt(0xFF0000),
		Color4:          helpers.RGBFromInt(0xFFA500),
		Color5:          helpers.RGBFromInt(0xFFFF00),
		Color6:          helpers.RGBFromInt(0x00FF00),
		Color7:          helpers.RGBFromInt(0x00BFFF),
		Color8:          helpers.RGBFromInt(0x0000FF),
		Color9:          helpers.RGBFromInt(0xFF00FF),
	}
}
