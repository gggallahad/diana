package handler

import "github.com/gggallahad/gui"

func (h *handler) redraw(ctx *gui.Context) {
	h.clear(ctx)

	h.drawTimerInfo(ctx)
	h.drawCurrentDirectoryInfo(ctx)
	h.drawCurrentDirectoryEntriesInfo(ctx)
	h.drawCursor(ctx)

	h.flush(ctx)
}
