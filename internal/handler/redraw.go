package handler

import "github.com/gggallahad/gui"

func (h *handler) redraw(ctx *gui.Context) {
	h.drawCurrentDirectoryInfo(ctx)
	h.drawCurrentDirectoryEntriesInfo(ctx)
	h.drawTimerInfo(ctx)
	h.DrawInterface(ctx, nil)
}
