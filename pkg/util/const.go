package util

import "github.com/gggallahad/gui"

const (
	TimerPositionX int    = 0
	TimerPositionY int    = 0
	TimerString    string = "Time: %.3f seconds"

	CurrentDirectoryNamePositionX int    = 3
	CurrentDirectoryNamePositionY int    = 2
	CurrentDirectoryStatPositionX int    = 0
	CurrentDirectoryStatPositionY int    = 2
	CurrentDirectoryNameString    string = "Directory name: %s"
	CurrentDirectoryStatString    string = "Total size: %s   Total entry count: %d   Done:%t"

	CurrentDirectoryEntryNamePositionX       int    = 10
	CurrentDirectoryEntryNamePositionY       int    = 4
	CurrentDirectoryEntryStatPositionX       int    = 0
	CurrentDirectoryEntryStatPositionY       int    = 4
	CurrentDirectoryEntryDirectoryNameString string = "%s "
	CurrentDirectoryEntryFileNameString      string = "%s "
	CurrentDirectoryEntryDirectoryStatString string = "Total size: %s   Total entry count: %d   Done:%t"
	CurrentDirectoryEntryFileStatString      string = "Size: %s"

	SizeString string = "%.1f %s"
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
