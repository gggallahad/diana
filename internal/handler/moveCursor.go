package handler

import (
	"github.com/gggallahad/diana/pkg/util"
)

func (h *handler) moveCursor(cursorPositionOffsetY int) {
	h.currentCursorPositionY += cursorPositionOffsetY

	currentDirectory := util.GetCurrentDirectory(h.previousDirectories)
	maxCursorY := len(currentDirectory.Entries) - 1

	if h.currentCursorPositionY < 0 {
		h.currentCursorPositionY = maxCursorY
	}

	if h.currentCursorPositionY > maxCursorY {
		h.currentCursorPositionY = 0
	}
}
