package models

import (
	"strings"
)

type History struct {
	Keys            string
	RunesBuffer     []rune
	ClearAfterMatch bool
}

const EmptyHistory = "                "

func NewHistory() *History {
	return &History{
		Keys:        EmptyHistory,
		RunesBuffer: []rune{},
	}
}

func (h *History) Clear() {
	h.Keys = EmptyHistory
}

func (h *History) Pressed(keys string) bool {
	pressed := strings.HasSuffix(h.Keys, keys)
	if pressed && h.ClearAfterMatch == true {
		h.Clear()
	}
	return pressed
}
