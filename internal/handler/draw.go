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
			h.ClearInterface(ctx, nil)

			h.drawTimerInfo(ctx)

			// if !h.isCurrentDirectoryUpToDate {
			h.drawCurrentDirectoryInfo(ctx)
			h.drawCurrentDirectoryEntriesInfo(ctx)
			// }

			// h.isCurrentDirectoryUpToDate = true

			h.drawCursor(ctx)

			h.DrawInterface(ctx, nil)

			break DrawLoop

		case <-h.drawTick:
			h.ClearInterface(ctx, nil)

			h.drawTimerInfo(ctx)

			// if !h.isCurrentDirectoryUpToDate {
			// h.checkCurrentDirectoryStatus()
			h.drawCurrentDirectoryInfo(ctx)
			h.drawCurrentDirectoryEntriesInfo(ctx)
			// }

			h.drawCursor(ctx)

			h.DrawInterface(ctx, nil)
		}
	}
}
