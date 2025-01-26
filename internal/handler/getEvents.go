package handler

import (
	"time"

	"github.com/gggallahad/gui"
)

func (h *handler) GetEvents(ctx *gui.Context) {
	for eventType := range h.eventChan {
		switch eventType.(type) {

		// case *model.EventUpdateEntryCount:

		// case *model.EventUpdateSize:

		}
	}

	h.endTime = time.Now()
	h.stopDraw <- struct{}{}
}
