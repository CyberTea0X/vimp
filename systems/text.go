package systems

import (
	"vimp/helpers"
	"vimp/models"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.design/x/clipboard"
)

func InputToText(a *models.App) {
	if a.GetMode() != models.ModeCommand {
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyV) {
		bytes := clipboard.Read(clipboard.FmtText)
		if len(bytes) > 128 {
			return
		}
		a.Text += string(bytes)
		return
	}
	if helpers.RepeatingKeyPressed(ebiten.KeyBackspace) {
		if len(a.Text) >= 1 {
			text := a.Text
			a.Text = text[:len(text)-1]
		}
		return
	}
	a.Text += string(a.History.RunesBuffer)
}
