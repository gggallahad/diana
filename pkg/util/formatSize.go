package util

import (
	"fmt"

	"github.com/gggallahad/diana/internal/model"
)

func FormatSize(format model.SizeFormat, size int64) string {
	switch format {

	case model.SizeFormatDynamic:
		return formatDynamic(size)

	case model.SizeFormatTerabytes:
		return formatTerabytes(size)

	case model.SizeFormatGigabytes:
		return formatGigabytes(size)

	case model.SizeFormatMegabytes:
		return formatMegabytes(size)

	case model.SizeFormatKilobytes:
		return formatKilobytes(size)

	case model.SizeFormatBytes:
		return formatBytes(size)

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

func formatTerabytes(size int64) string {
	terrabyteSize := float64(size) / float64(SizeTerrabyte)
	return fmt.Sprintf(SizeString, terrabyteSize, SizeFormatTerrabytesString)
}

func formatGigabytes(size int64) string {
	gigabyteSize := float64(size) / float64(SizeGigabyte)
	return fmt.Sprintf(SizeString, gigabyteSize, SizeFormatGigabytesString)
}

func formatMegabytes(size int64) string {
	megabyteSize := float64(size) / float64(SizeMegabyte)
	return fmt.Sprintf(SizeString, megabyteSize, SizeFormatMegabytesString)
}

func formatKilobytes(size int64) string {
	kilobyteSize := float64(size) / float64(SizeKilobyte)
	return fmt.Sprintf(SizeString, kilobyteSize, SizeFormatKilobytesString)
}

func formatBytes(size int64) string {
	return fmt.Sprintf(SizeString, float64(size), SizeFormatBytesString)
}
