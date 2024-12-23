package models

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Selection struct {
	Start image.Point
	End   image.Point
	Image *ebiten.Image
}
