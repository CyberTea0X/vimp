package models

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type ColorPicker struct {
	Image *ebiten.Image
	Draw  bool
}

func randomColor() color.Color {
	r := uint8(rand.Intn(255))
	g := uint8(rand.Intn(255))
	b := uint8(rand.Intn(255))
	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: 255,
	}
}

func (c *ColorPicker) Show(width, height int) {
	c.Draw = true
	c.Image = ebiten.NewImage(width, height)
}

func (c *ColorPicker) Resize(width, height int) {
	if c.Draw == false {
		return
	}
	i := ebiten.NewImage(width, height)
	i.DrawImage(c.Image, nil)
	c.Image = i
}

func (c *ColorPicker) FillRandomColor() {
	for i := 0; i < c.Image.Bounds().Dx(); i++ {
		for j := 0; j < c.Image.Bounds().Dy(); j++ {
			c.Image.Set(i, j, randomColor())
		}
	}
}
