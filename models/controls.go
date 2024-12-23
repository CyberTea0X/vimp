package models

//
// import (
// 	"github.com/hajimehoshi/ebiten/v2"
// 	"github.com/hajimehoshi/ebiten/v2/inpututil"
// )
//
// func AddCursorControls(k *Keymanager) {
// 	// Cursor controls
// 	k.AddHandler("j", ActionCursorDown)
// 	k.AddHandler("k", ActionCursorUp)
// 	k.AddHandler("h", ActionCursorLeft)
// 	k.AddHandler("l", ActionCursorRight)
// 	k.AddSpecialHandler(func(k *Keymanager) Action {
// 		if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
// 			return ActionCursorDown
// 		}
// 		if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
// 			return ActionCursorUp
// 		}
// 		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
// 			return ActionCursorLeft
// 		}
// 		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
// 			return ActionCursorRight
// 		}
// 		if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
// 			return ActionCursorDownStart
// 		}
// 		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
// 			return ActionCursorUpStart
// 		}
// 		if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
// 			return ActionCursorLeftStart
// 		}
// 		if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
// 			return ActionCursorRightStart
// 		}
// 		return ActionNone
// 	})
// }
//
// func AddNormalModeControls(k *Keymanager) {
// 	k.AddHandlers([]struct {
// 		key  string
// 		move Action
// 	}{
// 		{key: "cp", move: ActionColorPicker},
// 		{key: "ZZ", move: ActionLeaveAndSave},
// 		{key: "gg", move: ActionSetX0},
// 		{key: "G", move: ActionSetXLast},
// 		{key: "0", move: ActionSetY0},
// 		{key: "$", move: ActionSetYLast},
// 		{key: "ff", move: ActionFill},
// 		{key: "x", move: ActionDeletePixel},
// 		{key: "yp", move: ActionYankAsPrimaryColor},
// 		{key: "ys", move: ActionYankAsSecondaryColor},
// 		{key: "yy", move: ActionYankAsCurrentColor},
// 		{key: "p", move: ActionFillPixelPrimary},
// 		{key: "s", move: ActionFillPixelSecondary},
// 		{key: "1", move: ActionFillPixel1},
// 		{key: "2", move: ActionFillPixel2},
// 		{key: "3", move: ActionFillPixel3},
// 		{key: "4", move: ActionFillPixel4},
// 		{key: "5", move: ActionFillPixel5},
// 		{key: "6", move: ActionFillPixel6},
// 		{key: "7", move: ActionFillPixel7},
// 		{key: "8", move: ActionFillPixel8},
// 		{key: "9", move: ActionFillPixel9},
// 	})
// 	k.AddSpecialHandler(func(k *Keymanager) Action {
// 		if ebiten.IsKeyPressed(ebiten.KeySpace) {
// 			return ActionFillPixel
// 		}
// 		return ActionNone
// 	})
// }
