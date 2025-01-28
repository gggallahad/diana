package handler

import (
	"github.com/gggallahad/gui"
)

func (h *handler) Move(ctx *gui.Context, eventType gui.Event) {
	switch event := eventType.(type) {
	case *gui.EventKey:
		if event.Symbol == 'w' {
			h.moveCursor(-1)
		}
		if event.Symbol == 's' {
			h.moveCursor(1)
		}

		if event.Symbol == 'i' {
			h.stepInDirectory()
		}
		if event.Symbol == 'o' {
			h.stepOutDirectory()
		}
	}

	h.redraw(ctx)
}
