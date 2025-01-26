package util

func IsExecutable(fileMode uint32) bool {
	return fileMode&ExecutableMask != 0
}
