package handler

import (
	"fmt"
	"time"

	"github.com/gggallahad/diana/pkg/util"
	"github.com/gggallahad/gui"
)

func (h *handler) drawTimerInfo(ctx *gui.Context) {
	h.drawMutex.Lock()

	ctx.ClearRow(util.TimerPositionY)

	var timer time.Duration
	if h.mainDirectory.IsDone {
		timer = h.endTime.Sub(h.startTime)
	} else {
		timer = time.Since(h.startTime)
	}

	timerText := fmt.Sprintf(util.TimerString, timer.Seconds())

	viewSizeX, _ := ctx.ViewSize()

	timerInfoOffset := viewSizeX - 1 - util.TimerPositionX - len(timerText)
	if timerInfoOffset < 0 {
		timerInfoOffset = 0
	}

	ctx.SetText(timerInfoOffset, util.TimerPositionY, timerText, gui.DefaultColor, gui.DefaultColor)

	h.drawMutex.Unlock()
}
