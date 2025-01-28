package handler

import (
	"github.com/gggallahad/diana/pkg/util"
	"github.com/gggallahad/gui"
)

func (h *handler) drawCursor(ctx *gui.Context) {
	h.drawMutex.Lock()
	defer h.drawMutex.Unlock()

	cursorPositionY := h.currentCursorPositionY + util.CursorPositionY

	viewSizeX, _ := ctx.ViewSize()

	rightCursorOffsetX := viewSizeX - 1 - util.RightCursorPositionX - len(util.RightCursor)
	if rightCursorOffsetX < 0 {
		rightCursorOffsetX = 0
	}

	ctx.SetText(util.LeftCursorPositionX, cursorPositionY, util.LeftCursor, util.CursorColor, gui.DefaultColor)
	ctx.SetText(rightCursorOffsetX, cursorPositionY, util.RightCursor, util.CursorColor, gui.DefaultColor)
}
