package handler

func (h *handler) checkCurrentDirectoryStatus() {
	if h.currentDirectory.IsDone {
		h.isCurrentDirectoryUpToDate = true
	}
}
