package handler

import "github.com/gggallahad/gui"

func (h *handler) flush(ctx *gui.Context) {
	h.drawMutex.Lock()
	defer h.drawMutex.Unlock()

	err := ctx.Flush()
	if err != nil {
		return
	}
}
