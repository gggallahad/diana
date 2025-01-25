package util

var (
	Path            string
	PathFlagName    string = "p"
	DefaultPath     string = "."
	PathDescription string = "path to directory"

	FPS            int64
	FPSFlagName    string = "fps"
	DefaultFPS     int64  = 50
	FPSDescription string = "number of frames per second while directory stats are being calculated"
)
