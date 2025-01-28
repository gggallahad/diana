package util

import "github.com/gggallahad/gui"

const (
	TimerPositionX int    = 0
	TimerPositionY int    = 0
	TimerString    string = "Time: %.3f seconds"

	CurrentDirectoryNamePositionX int    = 7
	CurrentDirectoryNamePositionY int    = 2
	CurrentDirectoryStatPositionX int    = 7
	CurrentDirectoryStatPositionY int    = 2
	CurrentDirectoryNameString    string = "Directory name: %s"
	CurrentDirectoryStatString    string = "Total size: %s   Total entry count: %d   Done:%t"

	DirectorySeparator string = "-"

	CurrentDirectoryEntryNamePositionX       int    = 7
	CurrentDirectoryEntryNamePositionY       int    = 4
	CurrentDirectoryEntryStatPositionX       int    = 7
	CurrentDirectoryEntryStatPositionY       int    = 4
	CurrentDirectoryEntryDirectoryNameString string = "%s "
	CurrentDirectoryEntryFileNameString      string = "%s "
	CurrentDirectoryEntryDirectoryStatString string = "Total size: %s   Total entry count: %d   Done:%t"
	CurrentDirectoryEntryFileStatString      string = "Size: %s"

	SizeString string = "%.1f %s"

	LeftCursorPositionX  int = 2
	RightCursorPositionX int = 2
	CursorPositionY      int = 4
)

const (
	LeftCursor  string = "-->"
	RightCursor string = "<--"
)

const (
	SizeFormatBytesString      string = "b"
	SizeFormatKilobytesString  string = "kb"
	SizeFormatMegabytesString  string = "mb"
	SizeFormatGigabytesString  string = "gb"
	SizeFormatTerrabytesString string = "tb"
)

const (
	SizeKilobyte  int64 = 1024
	SizeMegabyte  int64 = 1024 * 1024
	SizeGigabyte  int64 = 1024 * 1024 * 1024
	SizeTerrabyte int64 = 1024 * 1024 * 1024 * 1024
)

var (
	CursorColor gui.Color = gui.Color{
		R: 162,
		G: 97,
		B: 208,
	}

	DirectoryColor gui.Color = gui.Color{
		R: 20,
		G: 75,
		B: 190,
	}

	ExecutableColor gui.Color = gui.Color{
		R: 38,
		G: 162,
		B: 105,
	}
)

const (
	ExecutableMask uint32 = 0111
)
