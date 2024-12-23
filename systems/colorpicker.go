package systems

import (
	"vimp/models"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawColorpicker(a *models.App) {
	if a.ColorPicker.Draw == false {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(a.Canvas.Scale, a.Canvas.Scale)
	a.Canvas.ImageOver.DrawImage(a.ColorPicker.Image, op)
}
