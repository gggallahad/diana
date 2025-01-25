package handler

import (
	"github.com/gggallahad/gui"
)

func (h *handler) Init(ctx *gui.Context) {
	err := ctx.Clear()
	if err != nil {
		return
	}

	h.drawCurrentDirectoryInfo(ctx)
	h.drawInterface(ctx)
}
