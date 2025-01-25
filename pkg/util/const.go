package util

const (
	TimerPositionX int    = 0
	TimerPositionY int    = 0
	TimerString    string = "Time: %.3f seconds"

	CurrentDirectoryInfoPositionX int    = 3
	CurrentDirectoryInfoPositionY int    = 1
	CurrentDirectoryNameString    string = "Directory name: %s"
	CurrentDirectoryStatString    string = "Total size: %s   Total entry count: %d   Done:%t"

	CurrentDirectoryEntriesPositionX  int    = 10
	CurrentDirectoryEntriesPositionY  int    = 2
	CurrentDirectoryEntriesNameString string = "Directory name: %s"
	CurrentDirectoryEntriesStatString string = "Total size: %s   Total entry count: %d   Done:%t"

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
