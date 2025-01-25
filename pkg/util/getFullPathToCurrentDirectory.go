package util

import (
	"path/filepath"

	"github.com/gggallahad/diana/internal/model"
)

func GetFullPathToCurrentDirectory(mainDirectory *model.EntryDirectory, pastDirectories []int) (string, error) {
	fullPath := Path
	currentDirectory := mainDirectory

	for i := range pastDirectories {
		if len(currentDirectory.Entries)-1 < pastDirectories[i] {
			return "", ErrNoEntry
		}

		entry := currentDirectory.Entries[pastDirectories[i]]
		nextDirectory, ok := entry.(*model.EntryDirectory)
		if !ok {
			return "", ErrEntryNotADirectory
		}

		fullPath = filepath.Join(fullPath, nextDirectory.Name)
		currentDirectory = nextDirectory
	}

	return fullPath, nil
}
