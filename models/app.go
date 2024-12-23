package models

import (
	"image/color"
	"image/png"
	"log"
	"os"
	"vimp/helpers"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.design/x/clipboard"
)

type AppMode int

const (
	ModeNormal AppMode = iota
	ModeVisual
	ModeCommand
)

type App struct {
	InfoBar         Bar
	CommandBar      Bar
	FilePath        string
	Selection       Selection
	DrawFunc        func(a *App, s *ebiten.Image)
	UpdateFunc      func(a *App) error
	Cursor          Cursor
	mode            AppMode
	previousMode    AppMode
	Palette         Palette
	Fonts           Fonts
	BackgroundColor color.Color
	Canvas          Canvas
	ColorPicker     ColorPicker
	History         History
	Text            string
}

type NewAppOpts struct {
	FilePath string
}

func NewApp(opts NewAppOpts) *App {
	if err := clipboard.Init(); err != nil {
		log.Fatal(err)
	}
	a := &App{
		BackgroundColor: helpers.RGBFromInt(0xFFFFFF),
		InfoBar:         *NewBar(256, 18, helpers.RGBFromInt(0xFFBBAA)),
		CommandBar:      *NewBar(256, 18, helpers.RGBFromInt(0x00AAFF)),
		DrawFunc:        func(a *App, s *ebiten.Image) {},
		UpdateFunc:      func(a *App) error { return nil },
		Cursor:          Cursor{X: 0, Y: 0},
		History:         *NewHistory(),
		Palette:         *DefaultPalette(),
	}
	var err error
	if opts.FilePath == "" {
		err = a.LoadFile("image.png")
	} else {
		err = a.LoadFile(opts.FilePath)
	}
	if err != nil {
		log.Fatal(err)
	}
	return a
}

func (a *App) GetMode() AppMode {
	return a.mode
}

func (a *App) SetMode(mode AppMode) {
	a.previousMode = a.mode
	a.mode = mode
}

func (a *App) GetPreviousMode() AppMode {
	return a.previousMode
}

func (a *App) LoadFile(path string) error {
	image, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			image = ebiten.NewImage(32, 32)
			image.Fill(a.BackgroundColor)
		} else {
			return err
		}
	}
	a.FilePath = path
	a.Canvas.ChangeCanvas(image)
	a.ColorPicker.Resize(a.Canvas.Width(), a.Canvas.Height())
	return nil
}

func (a *App) SaveToFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	if err := png.Encode(file, a.Canvas.Image); err != nil {
		return err
	}
	return nil
}

func (g *App) Update() error {
	return g.UpdateFunc(g)
}

func (a *App) Draw(screen *ebiten.Image) {
	a.DrawFunc(a, screen)
}

func (g *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	w := outsideWidth/4 + g.InfoBar.Height
	h := outsideHeight/4 + g.InfoBar.Height
	g.InfoBar.Width = w
	g.CommandBar.Width = w
	return w, h
}
