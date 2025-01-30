package handler

import (
	"github.com/gggallahad/gui"
)

func (h *handler) Init(ctx *gui.Context) {
	h.clear(ctx)

	h.drawCurrentDirectoryInfo(ctx)

	h.flush(ctx)
}
