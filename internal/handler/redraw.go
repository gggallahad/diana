package handler

import "github.com/gggallahad/gui"

func (h *handler) redraw(ctx *gui.Context) {
	h.ClearInterface(ctx, nil)

	h.drawTimerInfo(ctx)
	h.drawCurrentDirectoryInfo(ctx)
	h.drawCurrentDirectoryEntriesInfo(ctx)
	h.drawCursor(ctx)

	h.DrawInterface(ctx, nil)
}
