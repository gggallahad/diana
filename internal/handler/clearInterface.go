package handler

import "github.com/gggallahad/gui"

func (h *handler) ClearInterface(ctx *gui.Context, eventType gui.Event) {
	h.drawMutex.Lock()
	defer h.drawMutex.Unlock()

	err := ctx.Clear()
	if err != nil {
		return
	}
}
