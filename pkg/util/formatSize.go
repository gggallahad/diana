package util

import (
	"fmt"

	"github.com/gggallahad/diana/internal/model"
)

func FormatSize(format model.SizeFormat, size int64) string {
	switch format {

	case model.SizeFormatDynamic:
		return formatDynamic(size)

	case model.SizeFormatBytes:

	case model.SizeFormatKilobytes:

	case model.SizeFormatMegabytes:

	case model.SizeFormatGigabytes:

	case model.SizeFormatTerrabytes:

	}

	// unreachable
	return ""
}

func formatDynamic(size int64) string {
	switch {

	case size >= SizeTerrabyte:
		terrabyteSize := float64(size) / float64(SizeTerrabyte)
		return fmt.Sprintf(SizeString, terrabyteSize, SizeFormatTerrabytesString)

	case size >= SizeGigabyte:
		gigabyteSize := float64(size) / float64(SizeGigabyte)
		return fmt.Sprintf(SizeString, gigabyteSize, SizeFormatGigabytesString)

	case size >= SizeMegabyte:
		megabyteSize := float64(size) / float64(SizeMegabyte)
		return fmt.Sprintf(SizeString, megabyteSize, SizeFormatMegabytesString)

	case size >= SizeKilobyte:
		kilobyteSize := float64(size) / float64(SizeKilobyte)
		return fmt.Sprintf(SizeString, kilobyteSize, SizeFormatKilobytesString)

	default:
		return fmt.Sprintf(SizeString, float64(size), SizeFormatBytesString)
	}
}
