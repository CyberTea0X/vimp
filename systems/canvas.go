package systems

import (
	"vimp/helpers"
	"vimp/models"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawCanvas(a *models.App, s *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(a.Canvas.Scale, a.Canvas.Scale)
	op.GeoM.Translate(float64(a.InfoBar.Height)/2, float64(a.InfoBar.Height))
	s.DrawImage(a.Canvas.Image, op)
}

func DrawCanvasOver(a *models.App, s *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(a.InfoBar.Height)/2, float64(a.InfoBar.Height))
	s.DrawImage(a.Canvas.ImageOver, op)
	a.Canvas.ImageOver.Clear()
}

func DrawCanvasBorders(a *models.App, s *ebiten.Image) {
	vector.StrokeRect(
		s,
		float32(a.InfoBar.Height/2),
		float32(a.InfoBar.Height),
		float32(float64(a.Canvas.Image.Bounds().Dx())*a.Canvas.Scale+1),
		float32(float64(a.Canvas.Image.Bounds().Dy())*a.Canvas.Scale+1),
		1,
		helpers.RGBFromInt(0xACAAAA),
		false,
	)
}
