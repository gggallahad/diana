package handler

import (
	"github.com/gggallahad/gui"
)

func (h *handler) Init(ctx *gui.Context) {
	h.ClearInterface(ctx, nil)

	// h.drawTimerInfo(ctx)
	h.drawCurrentDirectoryInfo(ctx)

	h.DrawInterface(ctx, nil)
}
