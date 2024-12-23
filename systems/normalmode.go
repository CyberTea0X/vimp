package systems

import (
	"image/color"
	"vimp/helpers"
	"vimp/models"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func HandleNormalMode(a *models.App, i *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		a.Text = ""
		a.SetMode(models.ModeNormal)
		return nil
	}
	if a.GetMode() != models.ModeNormal {
		return nil
	}
	cursorX := float32(a.Cursor.X)
	cursorY := float32(a.Cursor.Y)
	var pixelColor color.Color
	if a.ColorPicker.Draw {
		pixelColor = a.ColorPicker.Image.At(int(a.Cursor.X), int(a.Cursor.Y))
	} else {
		pixelColor = a.Canvas.Image.At(int(a.Cursor.X), int(a.Cursor.Y))
	}
	a.History.ClearAfterMatch = true
	switch {
	case a.History.Pressed("cp"):
		if a.ColorPicker.Draw {
			a.ColorPicker.Draw = false
		} else {
			a.ColorPicker.Show(a.Canvas.Width(), a.Canvas.Height())
			a.ColorPicker.FillRandomColor()
		}
	case a.History.Pressed("ZZ"):
		a.SaveToFile(a.FilePath)
		return ebiten.Termination
	case a.History.Pressed("gg"):
		a.Cursor.Y = 0
	case a.History.Pressed("ff"):
		helpers.Fill(a.Canvas.Image, a.Palette.GetCurrentColor(), int(a.Cursor.X), int(a.Cursor.Y), 0)
	case a.History.Pressed("yp"):
		a.Palette.PrimaryColor = pixelColor
	case a.History.Pressed("ys"):
		a.Palette.SecondaryColor = pixelColor
	case a.History.Pressed("yy"):
		a.Palette.SetCurrentColor(pixelColor)
	case a.History.Pressed("G"):
		a.Cursor.Y = float64(a.Canvas.Image.Bounds().Dy() - 1)
	case a.History.Pressed("0"):
		a.Cursor.X = 0
	case a.History.Pressed("$"):
		a.Cursor.X = float64(a.Canvas.Image.Bounds().Dx() - 1)
	case a.History.Pressed("1"):
		a.Palette.SetCurrentColor(a.Palette.PrimaryColor)
	case a.History.Pressed("2"):
		a.Palette.SetCurrentColor(a.Palette.SecondaryColor)
	case a.History.Pressed("3"):
		a.Palette.SetCurrentColor(a.Palette.Color3)
	case a.History.Pressed("4"):
		a.Palette.SetCurrentColor(a.Palette.Color4)
	case a.History.Pressed("5"):
		a.Palette.SetCurrentColor(a.Palette.Color5)
	case a.History.Pressed("6"):
		a.Palette.SetCurrentColor(a.Palette.Color6)
	case a.History.Pressed("7"):
		a.Palette.SetCurrentColor(a.Palette.Color7)
	case a.History.Pressed("8"):
		a.Palette.SetCurrentColor(a.Palette.Color8)
	case a.History.Pressed("9"):
		a.Palette.SetCurrentColor(a.Palette.Color9)
	case a.History.Pressed("p"):
		if a.Selection.Image == nil {
			break
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(a.Cursor.X, a.Cursor.Y)
		a.Canvas.Image.DrawImage(a.Selection.Image, op)
	}
	a.History.ClearAfterMatch = false
	if ebiten.IsKeyPressed(ebiten.KeyX) {
		i.Set(int(cursorX), int(cursorY), a.Palette.SecondaryColor)
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		vector.DrawFilledRect(i, float32(a.Cursor.X), float32(a.Cursor.Y), 1, 1, a.Palette.GetCurrentColor(), false)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		vector.DrawFilledRect(i, float32(a.Cursor.X), float32(a.Cursor.Y), 1, 1, a.Palette.SecondaryColor, false)
	}
	return nil
}
