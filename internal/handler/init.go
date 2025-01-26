package handler

import (
	"github.com/gggallahad/gui"
)

func (h *handler) Init(ctx *gui.Context) {
	h.drawMutex.Lock()

	err := ctx.Clear()
	if err != nil {
		return
	}

	h.drawMutex.Unlock()

	h.drawCurrentDirectoryInfo(ctx)
	h.DrawInterface(ctx, nil)
}
