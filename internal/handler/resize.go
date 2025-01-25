package handler

import "github.com/gggallahad/gui"

func (h *handler) Resize(ctx *gui.Context, eventType gui.Event) {
	switch eventType.(type) {

	case *gui.EventResize:
		h.redraw(ctx)
	}
}
