package handler

import (
	"fmt"

	"github.com/gggallahad/diana/internal/model"
	"github.com/gggallahad/diana/pkg/util"
	"github.com/gggallahad/gui"
)

func (h *handler) drawCurrentDirectoryEntriesInfo(ctx *gui.Context) {
	h.drawMutex.Lock()
	defer h.drawMutex.Unlock()

	currentDirectory := util.GetCurrentDirectory(h.previousDirectories)
	entries := currentDirectory.Entries

	util.SortEntries(h.sortFormat, entries)

	y := util.CurrentDirectoryEntryNamePositionY
	for i := range entries {

		switch entry := entries[i].(type) {

		case *model.EntryDirectory:
			directoryName := fmt.Sprintf(util.CurrentDirectoryEntryDirectoryNameString, entry.Name)

			directorySize := util.FormatSize(h.sizeFormat, entry.TotalSize)

			directoryStat := fmt.Sprintf(util.CurrentDirectoryEntryDirectoryStatString, directorySize, entry.TotalEntryCount, entry.IsDone)

			viewSizeX, _ := ctx.ViewSize()

			directoryStatOffset := viewSizeX - 1 - util.CurrentDirectoryEntryStatPositionX - len(directoryStat)
			if directoryStatOffset < 0 {
				directoryStatOffset = 0
			}

			ctx.SetText(util.CurrentDirectoryEntryNamePositionX, y, directoryName, util.DirectoryColor, gui.DefaultColor)

			ctx.SetText(directoryStatOffset, y, directoryStat, gui.DefaultColor, gui.DefaultColor)

		case *model.EntryFile:
			fileName := fmt.Sprintf(util.CurrentDirectoryEntryFileNameString, entry.Name)

			fileSize := util.FormatSize(h.sizeFormat, entry.Size)

			fileStat := fmt.Sprintf(util.CurrentDirectoryEntryFileStatString, fileSize)

			viewSizeX, _ := ctx.ViewSize()

			fileNameColor := gui.DefaultColor
			isExecutable := util.IsExecutable(entry.Mode)
			if isExecutable {
				fileNameColor = util.ExecutableColor
			}

			fileStatOffset := viewSizeX - 1 - util.CurrentDirectoryEntryStatPositionX - len(fileStat)
			if fileStatOffset < 0 {
				fileStatOffset = 0
			}

			ctx.SetText(util.CurrentDirectoryEntryNamePositionX, y, fileName, fileNameColor, gui.DefaultColor)

			ctx.SetText(fileStatOffset, y, fileStat, gui.DefaultColor, gui.DefaultColor)
		}

		y++
	}
}
