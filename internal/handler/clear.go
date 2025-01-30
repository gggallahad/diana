package handler

import "github.com/gggallahad/gui"

func (h *handler) clear(ctx *gui.Context) {
	h.drawMutex.Lock()
	defer h.drawMutex.Unlock()

	err := ctx.Clear()
	if err != nil {
		return
	}
}
