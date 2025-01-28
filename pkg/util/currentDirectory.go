package util

import (
	"path/filepath"

	"github.com/gggallahad/diana/internal/model"
)

func GetFullPathToCurrentDirectory(pastDirectories []*model.EntryDirectory) (string, error) {
	fullPath := Path
	// currentDirectory := pastDirectories[0]

	for i := 1; i < len(pastDirectories); i++ {
		fullPath = filepath.Join(fullPath, pastDirectories[i].Name)
		// currentDirectory = pastDirectories[i]
	}

	return fullPath, nil
}

func GetCurrentDirectory(pastDirectories []*model.EntryDirectory) *model.EntryDirectory {
	return pastDirectories[len(pastDirectories)-1]
}
