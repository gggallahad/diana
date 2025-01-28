package handler

func (h *handler) stepOutDirectory() {
	h.drawMutex.Lock()
	defer h.drawMutex.Unlock()

	lenPreviousDirectories := len(h.previousDirectories)
	if lenPreviousDirectories <= 1 {
		return
	}

	h.previousDirectories = h.previousDirectories[:lenPreviousDirectories-1]

	h.currentCursorPositionY = 0
}
