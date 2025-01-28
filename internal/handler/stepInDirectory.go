package handler

import (
	"github.com/gggallahad/diana/internal/model"
	"github.com/gggallahad/diana/pkg/util"
)

func (h *handler) stepInDirectory() {
	h.drawMutex.Lock()
	defer h.drawMutex.Unlock()

	currentDirectory := util.GetCurrentDirectory(h.previousDirectories)

	maxEntryIndex := len(currentDirectory.Entries) - 1
	if h.currentCursorPositionY > maxEntryIndex {
		return
	}

	entryToStep := currentDirectory.Entries[h.currentCursorPositionY]

	directoryToStep, ok := entryToStep.(*model.EntryDirectory)
	if !ok {
		return
	}

	h.previousDirectories = append(h.previousDirectories, directoryToStep)

	h.currentCursorPositionY = 0
}
