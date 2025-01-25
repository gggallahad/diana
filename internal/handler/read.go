package handler

import (
	"time"

	"github.com/gggallahad/diana/pkg/util"
	"github.com/gggallahad/gui"
)

func (h *handler) Read(ctx *gui.Context) {
	h.startTime = time.Now()

	h.eventWG.Add(1)
	go util.CloseChan(h.eventChan, &h.eventWG)

	go h.Service.Read(util.Path, h.mainDirectory, h.eventChan, &h.eventWG)
}
