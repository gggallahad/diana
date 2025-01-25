package handler

import (
	"fmt"
	"strings"

	"github.com/gggallahad/diana/pkg/util"
	"github.com/gggallahad/gui"
)

func (h *handler) drawCurrentDirectoryInfo(ctx *gui.Context) {
	h.drawMutex.Lock()

	ctx.ClearRow(util.CurrentDirectoryInfoPositionY)

	directoryFullPath, err := util.GetFullPathToCurrentDirectory(h.mainDirectory, h.pastDirectories)
	if err != nil {
		return
	}

	directorySize := util.FormatSize(h.sizeFormat, h.currentDirectory.TotalSize)

	directoryName := fmt.Sprintf(util.CurrentDirectoryNameString, directoryFullPath)
	directoryStat := fmt.Sprintf(util.CurrentDirectoryStatString, directorySize, h.currentDirectory.TotalEntryCount, h.currentDirectory.IsDone)

	viewSizeX, _ := ctx.ViewSize()
	repeatCount := viewSizeX - 1 - util.CurrentDirectoryInfoPositionX - len(directoryStat) - len(directoryName)
	if repeatCount < 1 {
		repeatCount = 1
	}
	spaces := strings.Repeat(" ", repeatCount)

	directoryInfo := fmt.Sprintf("%s%s%s", directoryName, spaces, directoryStat)
	ctx.SetText(util.CurrentDirectoryInfoPositionX, util.CurrentDirectoryInfoPositionY, directoryInfo, gui.DefaultColor, gui.DefaultColor)

	h.drawMutex.Unlock()
}
