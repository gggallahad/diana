package handler

import "github.com/gggallahad/gui"

func (h *handler) DrawInterface(ctx *gui.Context, eventType gui.Event) {
	h.drawMutex.Lock()
	defer h.drawMutex.Unlock()

	err := ctx.Flush()
	if err != nil {
		return
	}
}
