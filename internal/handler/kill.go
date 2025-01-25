package handler

import "github.com/gggallahad/gui"

func (h *handler) Kill(ctx *gui.Context, eventType gui.Event) {
	switch event := eventType.(type) {

	case *gui.EventKey:
		if event.Key == gui.KeyEsc || event.Symbol == 'q' {
			ctx.Abort()
			ctx.Kill()
		}
	}
}
