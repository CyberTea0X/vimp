package models

type Action int

const (
	ActionNone Action = iota
	// Cursor
	ActionCursorUp
	ActionCursorDown
	ActionCursorLeft
	ActionCursorRight
	ActionCursorUpStart
	ActionCursorDownStart
	ActionCursorLeftStart
	ActionCursorRightStart
	// Mode switch
	ActionLeaveMode
	ActionEnterCommandMode
	ActionCommandEntered
	// Normal mode
	ActionColorPicker
	ActionLeaveAndSave
	ActionSetX0
	ActionSetXLast
	ActionSetY0
	ActionSetYLast
	ActionFill
	ActionDeletePixel
	ActionYankAsPrimaryColor
	ActionYankAsSecondaryColor
	ActionYankAsCurrentColor
	ActionFillPixel
	ActionFillPixelPrimary
	ActionFillPixelSecondary
	ActionFillPixel1
	ActionFillPixel2
	ActionFillPixel3
	ActionFillPixel4
	ActionFillPixel5
	ActionFillPixel6
	ActionFillPixel7
	ActionFillPixel8
	ActionFillPixel9
)
