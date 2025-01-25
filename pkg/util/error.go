package util

import "errors"

var (
	ErrReadDirectory      error = errors.New("read directory error")
	ErrGetFileInfo        error = errors.New("get file info error")
	ErrNoEntry            error = errors.New("no entry error")
	ErrEntryNotADirectory error = errors.New("entry not a directory error")
)
