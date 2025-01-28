package handler

import (
	"fmt"
	"strings"

	"github.com/gggallahad/diana/pkg/util"
	"github.com/gggallahad/gui"
)

func (h *handler) drawCurrentDirectoryInfo(ctx *gui.Context) {
	h.drawMutex.Lock()
	defer h.drawMutex.Unlock()

	currentDirectory := util.GetCurrentDirectory(h.previousDirectories)
	directoryFullPath, err := util.GetFullPathToCurrentDirectory(h.previousDirectories)
	if err != nil {
		return
	}

	directoryName := fmt.Sprintf(util.CurrentDirectoryNameString, directoryFullPath)

	directorySize := util.FormatSize(h.sizeFormat, currentDirectory.TotalSize)

	directoryStat := fmt.Sprintf(util.CurrentDirectoryStatString, directorySize, currentDirectory.TotalEntryCount, currentDirectory.IsDone)

	viewSizeX, _ := ctx.ViewSize()

	directoryStatOffsetX := viewSizeX - 1 - util.CurrentDirectoryStatPositionX - len(directoryStat)
	if directoryStatOffsetX < 0 {
		directoryStatOffsetX = 0
	}

	ctx.SetText(util.CurrentDirectoryNamePositionX, util.CurrentDirectoryNamePositionY, directoryName, gui.DefaultColor, gui.DefaultColor)

	ctx.SetText(directoryStatOffsetX, util.CurrentDirectoryStatPositionY, directoryStat, gui.DefaultColor, gui.DefaultColor)

	directorySeparatorPositionXStart := util.CurrentDirectoryNamePositionX
	directorySeparatorPositionXEnd := directoryStatOffsetX + len(directoryStat)

	directorySeparatorCount := directorySeparatorPositionXEnd - directorySeparatorPositionXStart
	directorySeparator := strings.Repeat(util.DirectorySeparator, directorySeparatorCount)

	ctx.SetText(util.CurrentDirectoryNamePositionX, util.CurrentDirectoryNamePositionY+1, directorySeparator, gui.DefaultColor, gui.DefaultColor)
}
