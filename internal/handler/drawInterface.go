package handler

import "github.com/gggallahad/gui"

func (h *handler) drawInterface(ctx *gui.Context) {
	h.drawMutex.Lock()

	err := ctx.Flush()
	if err != nil {
		return
	}

	h.drawMutex.Unlock()
}
