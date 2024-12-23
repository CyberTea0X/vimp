package systems

import (
	"strings"
	"vimp/models"

	"github.com/hajimehoshi/ebiten/v2"
)

func UpdateHistory(a *models.History) {
	a.RunesBuffer = ebiten.AppendInputChars(a.RunesBuffer[:0])
	if len(a.RunesBuffer) != 0 {
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			a.Keys = a.Keys[len(a.RunesBuffer):] + strings.ToUpper(string(a.RunesBuffer))
		} else {
			a.Keys = a.Keys[len(a.RunesBuffer):] + string(a.RunesBuffer)
		}
	}
}
