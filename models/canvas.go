package models

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Canvas struct {
	Image     *ebiten.Image
	ImageOver *ebiten.Image
	Scale     float64
}

func (c *Canvas) ChangeCanvas(i *ebiten.Image) {
	c.Image = i
	c.Scale = 256 / math.Max(float64(c.Image.Bounds().Dx()), float64(c.Image.Bounds().Dy()))
	c.ImageOver = ebiten.NewImage(c.Image.Bounds().Dx()*int(c.Scale), c.Image.Bounds().Dy()*int(c.Scale))
}

func (c *Canvas) ResizeCanvas(newScale float64) {
	width := float64(c.Image.Bounds().Dx()) * newScale
	height := float64(c.Image.Bounds().Dy()) * newScale
	i := ebiten.NewImage(int(width), int(height))
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(newScale, newScale)
	i.DrawImage(c.Image, op)
	c.ChangeCanvas(i)
}

func (c *Canvas) Width() int {
	return c.Image.Bounds().Dx()
}

func (c *Canvas) Height() int {
	return c.Image.Bounds().Dy()
}
