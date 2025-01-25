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
			if !h.isCurrentDirectoryUpToDate {
				h.drawCurrentDirectoryInfo(ctx)
			}

			h.isCurrentDirectoryUpToDate = true

			h.drawTimerInfo(ctx)

			h.drawInterface(ctx)

			break DrawLoop

		case <-h.drawTick:
			if !h.isCurrentDirectoryUpToDate {
				h.checkCurrentDirectoryStatus()
				h.drawCurrentDirectoryInfo(ctx)
			}

			h.drawTimerInfo(ctx)

			h.drawInterface(ctx)
		}
	}
}
