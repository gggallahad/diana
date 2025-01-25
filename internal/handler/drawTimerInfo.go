package handler

import (
	"fmt"
	"strings"
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
	repeatCount := viewSizeX - 1 - util.TimerPositionX - len(timerText)
	if repeatCount < 1 {
		repeatCount = 1
	}
	spaces := strings.Repeat(" ", repeatCount)

	timerInfo := fmt.Sprintf("%s%s", spaces, timerText)
	ctx.SetText(util.TimerPositionX, util.TimerPositionY, timerInfo, gui.DefaultColor, gui.DefaultColor)

	h.drawMutex.Unlock()
}
