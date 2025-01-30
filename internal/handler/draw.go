package handler

import (
	"time"

	"github.com/gggallahad/diana/pkg/util"
	"github.com/gggallahad/gui"
)

func (h *handler) Draw(ctx *gui.Context) {
	h.drawTick = time.Tick(time.Duration(time.Second / time.Duration(util.FPS)))

DrawLoop:
	for {
		select {

		case <-h.stopDraw:
			h.redraw(ctx)
			break DrawLoop

		case <-h.drawTick:
			h.redraw(ctx)
		}
	}
}
