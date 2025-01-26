package model

type (
	SizeFormat int
)

const (
	SizeFormatDynamic   SizeFormat = 0
	SizeFormatBytes     SizeFormat = 1
	SizeFormatKilobytes SizeFormat = 2
	SizeFormatMegabytes SizeFormat = 3
	SizeFormatGigabytes SizeFormat = 4
	SizeFormatTerabytes SizeFormat = 5
)
