package systems

import (
	"fmt"
	"math"
	"os"
	"strings"
	"vimp/helpers"
	"vimp/models"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func HandleCommandMode(a *models.App) error {
	if a.History.Pressed(":") {
		a.SetMode(models.ModeCommand)
		a.History.Clear()
		return nil
	}
	if a.GetMode() != models.ModeCommand {
		return nil
	}
	if !ebiten.IsKeyPressed(ebiten.KeyEnter) {
		return nil
	}
	parts := strings.SplitN(a.Text, " ", 2)
	if len(parts) == 0 {
		return nil
	}
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	command := parts[0]
	switch models.Command(command) {
	case models.CommandQuit:
		return ebiten.Termination
	case models.CommandQuitNoSave:
		return ebiten.Termination
	case models.CommandWrite:
		if len(parts) == 1 {
			a.SaveToFile(a.FilePath)
		} else if len(parts) == 2 {
			a.SaveToFile(parts[1])
		}
	case models.CommandColor:
		if len(parts) != 2 {
			break
		}
		color := parts[1]
		c, ok := models.HtmlColors[strings.ToUpper(parts[1])]
		if ok {
			a.Palette.SetCurrentColor(c)
			break
		}
		if !strings.HasPrefix(color, "0x") {
			color = "0x" + color
		}
		c = helpers.RGBAFromString(color)
		a.Palette.SetCurrentColor(c)
	case models.CommandOpen:
		if len(parts) != 2 {
			break
		}
		if err := a.LoadFile(parts[1]); err != nil {
			return err
		}
	case models.CommandClear:
		a.Canvas.Image.Fill(a.BackgroundColor)
	case models.CommandSize:
		if len(parts) != 2 {
			break
		}
		parts := strings.Split(parts[1], "x")
		if len(parts) != 2 {
			break
		}
		var w, h int
		fmt.Sscan(parts[0], &w)
		fmt.Sscan(parts[1], &h)
		canvas := ebiten.NewImage(w, h)
		canvas.DrawImage(a.Canvas.Image, nil)
		a.Canvas.ChangeCanvas(canvas)
		a.Cursor.X = 0
		a.Cursor.Y = 0
	case models.CommandRemove:
		if len(parts) == 1 {
			err := os.Remove(a.FilePath)
			if err != nil {
				return err
			}
			if err = a.LoadFile("image.png"); err != nil {
				return err
			}
			break
		}
		if len(parts) == 2 {
			err := os.Remove(parts[1])
			if err != nil {
				return err
			}
			if parts[1] == a.FilePath {
				err = a.LoadFile("image.png")
			}
			if err != nil {
				return err
			}

		}
	case models.CommandScale:
		if len(parts) != 2 {
			break
		}
		var canvasScale float64
		fmt.Sscan(parts[1], &canvasScale)
		a.Canvas.ResizeCanvas(canvasScale)
	case models.CommandCircle:
		if a.GetPreviousMode() != models.ModeVisual {
			break
		}
		x := a.Selection.Start.X
		y := a.Selection.Start.Y
		x1 := a.Selection.End.X
		y1 := a.Selection.End.Y
		vector.DrawFilledCircle(
			a.Canvas.Image,
			float32(x+x1)/2,
			float32(y+y1)/2,
			float32(math.Abs(float64(x-x1)))/2,
			a.Palette.GetCurrentColor(),
			false,
		)
	}
	a.Text = ""
	a.SetMode(a.GetPreviousMode())
	a.History.Clear()
	return nil
}
