package models

import (
	"os"
	"path"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	NormalFontSize = 16
	BigFontSize    = 24
)

type Fonts struct {
	FiraMono *text.GoTextFaceSource
}

func LoadFonts() (Fonts, error) {
	f := Fonts{}
	file, err := os.Open(path.Join(BasePath, "assets", "fonts", "FiraMono", "FiraMono-Regular.ttf"))
	if err != nil {
		return f, err
	}
	f.FiraMono, err = text.NewGoTextFaceSource(file)
	if err != nil {
		return f, err
	}
	return f, nil
}
