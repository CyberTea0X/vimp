package systems

import (
	"fmt"
	"image"
	"math"
	"path"
	"vimp/models"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawBar(a *models.App, s *ebiten.Image) {
	i := s.SubImage(image.Rect(0, 0, a.InfoBar.Width, a.InfoBar.Height)).(*ebiten.Image)
	i.Fill(a.InfoBar.Background)
	ebitenutil.DebugPrint(i, fmt.Sprintf(
		"%s | %dx%d | %d, %d",
		path.Base(a.FilePath),
		a.Canvas.Image.Bounds().Dx(),
		a.Canvas.Image.Bounds().Dy(),
		int(math.Round(a.Cursor.X)),
		int(math.Round(a.Cursor.Y)),
	))
	if a.GetMode() != models.ModeCommand {
		return
	}
	y0 := int(float64(a.Canvas.Image.Bounds().Dy())*a.Canvas.Scale + float64(a.InfoBar.Height))
	y1 := int(float64(a.Canvas.Image.Bounds().Dy()) * a.Canvas.Scale)
	i = s.SubImage(image.Rect(0, y0, a.CommandBar.Width, y1)).(*ebiten.Image)
	i.Fill(a.CommandBar.Background)
	ebitenutil.DebugPrintAt(i, fmt.Sprintf(":%s", a.Text), i.Bounds().Min.X, i.Bounds().Min.Y)
}
