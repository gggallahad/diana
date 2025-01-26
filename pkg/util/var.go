package util

import "github.com/gggallahad/diana/internal/model"

var (
	Path             string
	FPS              int64
	NumCPU           int
	SortFormatString string
	SortFormat       model.SortFormat
	SizeFormatString string
	SizeFormat       model.SizeFormat
)

const (
	PathFlagName    string = "p"
	DefaultPath     string = "."
	PathDescription string = "path to directory"

	FPSFlagName    string = "fps"
	DefaultFPS     int64  = 25
	FPSDescription string = "number of frames per second while directory stats are being calculated"

	NumCPUFlagName    string = "ncpu"
	DefaultNumCPU     int    = 0 // поставит runtime.NumCPU
	NumCPUDescription string = "max number of logical CPUs"

	SortFormatFlagName    = "sort"
	DefaultSortFormat     = "dsize"
	SortFormatDescription = "sort format for entries"

	SizeFormatFlagName    = "size"
	DefaultSizeFormat     = "dynamic"
	SizeFormatDescription = "size format"
)
