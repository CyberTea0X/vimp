package systems

import (
	"image"
	"math"
	"vimp/models"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func HandleVisualMode(a *models.App) {
	mode := a.GetMode()
	if mode == models.ModeNormal && inpututil.IsKeyJustPressed(ebiten.KeyV) {
		if mode == models.ModeVisual {
			a.SetMode(models.ModeNormal)
		} else {
			a.SetMode(models.ModeVisual)
			a.Selection.Start.X = int(math.Round(a.Cursor.X))
			a.Selection.Start.Y = int(math.Round(a.Cursor.Y))
			a.Selection.End.X = int(math.Round(a.Cursor.X))
			a.Selection.End.Y = int(math.Round(a.Cursor.Y))
		}
	}
	if a.GetMode() != models.ModeVisual {
		return
	}
	a.Selection.End.X = int(math.Round(a.Cursor.X))
	a.Selection.End.Y = int(math.Round(a.Cursor.Y))
	x := a.Selection.Start.X
	y := a.Selection.Start.Y
	x1 := a.Selection.End.X
	y1 := a.Selection.End.Y
	rect := image.Rect(x, y, x1, y1)
	a.History.ClearAfterMatch = true
	switch {
	case a.History.Pressed("y"):
		sub := a.Canvas.Image.SubImage(rect).(*ebiten.Image)
		a.Selection.Image = ebiten.NewImage(sub.Bounds().Dx(), sub.Bounds().Dy())
		a.Selection.Image.DrawImage(a.Canvas.Image.SubImage(rect).(*ebiten.Image), nil)
		a.SetMode(models.ModeNormal)
	case a.History.Pressed("d"):
		sub := a.Canvas.Image.SubImage(rect).(*ebiten.Image)
		a.Selection.Image = ebiten.NewImage(sub.Bounds().Dx(), sub.Bounds().Dy())
		a.Selection.Image.DrawImage(a.Canvas.Image.SubImage(rect).(*ebiten.Image), nil)
		sub.Fill(a.Palette.BackgroundColor)
		a.SetMode(models.ModeNormal)
	case a.History.Pressed("r1"):
		a.Canvas.Image.SubImage(rect).(*ebiten.Image).Fill(a.Palette.PrimaryColor)
		a.SetMode(models.ModeNormal)
	case a.History.Pressed("r2"):
		a.Canvas.Image.SubImage(rect).(*ebiten.Image).Fill(a.Palette.SecondaryColor)
		a.SetMode(models.ModeNormal)
	}
	a.History.ClearAfterMatch = false
}

func DrawSelection(a *models.App, i *ebiten.Image) {
	if a.GetMode() != models.ModeVisual && a.GetMode() != models.ModeCommand {
		return
	}
	x := a.Selection.Start.X
	y := a.Selection.Start.Y
	width := a.Selection.End.X - x
	height := a.Selection.End.Y - y
	vector.StrokeRect(
		i,
		float32(float64(x)*a.Canvas.Scale)+1,
		float32(float64(y)*a.Canvas.Scale)+1,
		float32(float64(width)*a.Canvas.Scale)-1,
		float32(float64(height)*a.Canvas.Scale)-1,
		1,
		a.Palette.GetCurrentColor(),
		false,
	)
}
