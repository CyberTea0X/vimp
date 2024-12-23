package helpers

import (
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func IsAnyKeyPressed(keys ...ebiten.Key) bool {
	for _, key := range keys {
		if ebiten.IsKeyPressed(key) {
			return true
		}
	}
	return false
}

func IsAnyKeyJustPressed(keys ...ebiten.Key) bool {
	for _, key := range keys {
		if inpututil.IsKeyJustPressed(key) {
			return true
		}
	}
	return false
}

func IsAllKeysJustPressed(keys ...ebiten.Key) bool {
	for _, key := range keys {
		if !inpututil.IsKeyJustPressed(key) {
			return false
		}
	}
	return true
}

func IsKeySeqPressed(runes, history []rune, seq string) bool {
	if !slices.Contains(runes, rune(seq[len(seq)-1])) {
		return false
	}
	for i := 0; i < len(seq); i++ {
		if history[i] != rune(seq[len(seq)-i-1]) {
			return false
		}
	}
	return true
}

func IsLowerKeyPressed(key ebiten.Key) bool {
	return ebiten.IsKeyPressed(key) && !ebiten.IsKeyPressed(ebiten.KeyShift)
}

func IsUpperKeyPressed(key ebiten.Key) bool {
	return ebiten.IsKeyPressed(key) && ebiten.IsKeyPressed(ebiten.KeyShift)
}

func KeyFromChar(char string) (ebiten.Key, bool) {
	switch char {
	case "a":
		return ebiten.KeyA, true
	case "b":
		return ebiten.KeyB, true
	case "c":
		return ebiten.KeyC, true
	case "d":
		return ebiten.KeyD, true
	case "e":
		return ebiten.KeyE, true
	case "f":
		return ebiten.KeyF, true
	case "g":
		return ebiten.KeyG, true
	case "h":
		return ebiten.KeyH, true
	case "i":
		return ebiten.KeyI, true
	case "j":
		return ebiten.KeyJ, true
	case "k":
		return ebiten.KeyK, true
	case "l":
		return ebiten.KeyL, true
	case "m":
		return ebiten.KeyM, true
	case "n":
		return ebiten.KeyN, true
	case "o":
		return ebiten.KeyO, true
	case "p":
		return ebiten.KeyP, true
	case "q":
		return ebiten.KeyQ, true
	case "r":
		return ebiten.KeyR, true
	case "s":
		return ebiten.KeyS, true
	case "t":
		return ebiten.KeyT, true
	case "u":
		return ebiten.KeyU, true
	case "v":
		return ebiten.KeyV, true
	case "w":
		return ebiten.KeyW, true
	case "x":
		return ebiten.KeyX, true
	case "y":
		return ebiten.KeyY, true
	case "z":
		return ebiten.KeyZ, true
	case "0":
		return ebiten.KeyDigit0, true
	case "1":
		return ebiten.KeyDigit1, true
	case "2":
		return ebiten.KeyDigit2, true
	case "3":
		return ebiten.KeyDigit3, true
	case "4":
		return ebiten.KeyDigit4, true
	case "5":
		return ebiten.KeyDigit5, true
	case "6":
		return ebiten.KeyDigit6, true
	case "7":
		return ebiten.KeyDigit7, true
	case "8":
		return ebiten.KeyDigit8, true
	case "9":
		return ebiten.KeyDigit9, true
	case " ":
		return ebiten.KeySpace, true
	}
	return ebiten.KeyDelete, false
}

func RepeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}
