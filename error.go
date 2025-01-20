package main

import "errors"

var (
	ErrReadDirectory error = errors.New("read directory error")
	ErrGetFileInfo   error = errors.New("get file info error")
)
