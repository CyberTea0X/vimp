package main

import (
	"errors"
	"image/color"
	"log"
	"vimp/models"
	"vimp/systems"

	"github.com/hajimehoshi/ebiten/v2"
)

const help = `
# NORMAL MODE:
hjkl or arrows - move
i - draw mode
v - visual mode
c - color mode
: - command
x - erase on cursor
space - use tool
? - help
zz - exit
ESC in any mode - back to normal mode

# DRAW MODE:
hjkl or arrows - move
p - primary color
s - secondary color
ESC - exit to normal mode

# COLOR MODE:
r - select red
g - select green
b - select blue
a - select alpha
`

const commands = `
:w  - write
:cd - change directory
:o <file>  - open file
:q - quit
`

func draw(a *models.App, s *ebiten.Image) {
	s.Fill(color.White)
	systems.DrawCanvas(a, s)
	systems.DrawCanvasOver(a, s)
	systems.DrawColorpicker(a)
	systems.DrawCursor(a, a.Canvas.ImageOver)
	systems.DrawSelection(a, a.Canvas.ImageOver)
	systems.DrawCanvasBorders(a, s)
	systems.DrawBar(a, s)
}

func update(a *models.App) error {
	systems.UpdateHistory(&a.History)
	systems.InputToText(a)
	systems.MoveCursor(a)
	err := systems.HandleNormalMode(a, a.Canvas.Image)
	if err != nil {
		return err
	}
	systems.HandleVisualMode(a)
	err = systems.HandleCommandMode(a)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ebiten.SetWindowSize(1024, 1040)
	ebiten.SetWindowTitle("PixelPaint")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	app := models.NewApp(models.NewAppOpts{})
	app.DrawFunc = draw
	app.UpdateFunc = update
	if err := ebiten.RunGame(app); err != nil && !errors.Is(err, ebiten.Termination) {
		log.Fatal(err)
	}
}
