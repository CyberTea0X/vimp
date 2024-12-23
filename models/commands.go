package models

type Command string

const (
	CommandQuit       Command = "q"
	CommandQuitNoSave Command = "q!"
	CommandWrite      Command = "w"
	CommandColor      Command = "color"
	CommandClear      Command = "clear"
	CommandOpen       Command = "o"
	CommandSize       Command = "size"
	CommandFill       Command = "fill"
	CommandRemove     Command = "rm"
	CommandScale      Command = "scale"
	CommandCircle     Command = "circle"
)
