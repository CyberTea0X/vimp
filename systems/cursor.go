package systems

import (
	"math"
	"vimp/helpers"
	"vimp/models"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func MoveCursor(a *models.App) {
	if a.GetMode() == models.ModeCommand {
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		a.Cursor.Y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		a.Cursor.Y += a.Cursor.Velocity
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyK) {
		a.Cursor.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		a.Cursor.Y -= a.Cursor.Velocity
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		a.Cursor.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyH) {
		a.Cursor.X -= a.Cursor.Velocity
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyL) {
		a.Cursor.X += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyL) {
		a.Cursor.X += a.Cursor.Velocity
	}

	if !helpers.IsAnyKeyPressed(
		ebiten.KeyJ,
		ebiten.KeyK,
		ebiten.KeyL,
		ebiten.KeyH,
		ebiten.KeyArrowDown,
		ebiten.KeyArrowUp,
		ebiten.KeyArrowRight,
		ebiten.KeyArrowLeft,
	) {
		a.Cursor.X = math.Round(a.Cursor.X)
		a.Cursor.Y = math.Round(a.Cursor.Y)
		a.Cursor.Velocity = 0.0
	} else {
		a.Cursor.Velocity += 0.01
	}
}

func DrawCursor(a *models.App, image *ebiten.Image) {
	vector.StrokeRect(
		image,
		float32(math.Round(a.Cursor.X)*a.Canvas.Scale)+1,
		float32(math.Round(a.Cursor.Y)*a.Canvas.Scale)+1,
		float32(a.Canvas.Scale)-1,
		float32(a.Canvas.Scale)-1,
		1,
		a.Palette.GetCurrentColor(),
		false,
	)
}
