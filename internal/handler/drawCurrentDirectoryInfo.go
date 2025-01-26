package handler

import (
	"fmt"

	"github.com/gggallahad/diana/pkg/util"
	"github.com/gggallahad/gui"
)

func (h *handler) drawCurrentDirectoryInfo(ctx *gui.Context) {
	h.drawMutex.Lock()

	ctx.ClearRow(util.CurrentDirectoryNamePositionY)

	directoryFullPath, err := util.GetFullPathToCurrentDirectory(h.mainDirectory, h.pastDirectories)
	if err != nil {
		return
	}

	directoryName := fmt.Sprintf(util.CurrentDirectoryNameString, directoryFullPath)

	directorySize := util.FormatSize(h.sizeFormat, h.currentDirectory.TotalSize)

	directoryStat := fmt.Sprintf(util.CurrentDirectoryStatString, directorySize, h.currentDirectory.TotalEntryCount, h.currentDirectory.IsDone)

	viewSizeX, _ := ctx.ViewSize()

	directoryStatOffset := viewSizeX - 1 - util.CurrentDirectoryStatPositionX - len(directoryStat)
	if directoryStatOffset < 0 {
		directoryStatOffset = 0
	}

	ctx.SetText(util.CurrentDirectoryNamePositionX, util.CurrentDirectoryNamePositionY, directoryName, gui.DefaultColor, gui.DefaultColor)

	ctx.SetText(directoryStatOffset, util.CurrentDirectoryStatPositionY, directoryStat, gui.DefaultColor, gui.DefaultColor)

	h.drawMutex.Unlock()
}
