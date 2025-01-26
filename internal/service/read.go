package service

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/gggallahad/diana/internal/model"
	"github.com/gggallahad/diana/pkg/util"
)

func (s *service) Read(path string, directory *model.EntryDirectory, eventChan chan<- model.Event, eventWG *sync.WaitGroup) {
	defer eventWG.Done()
	defer directory.Done()

	dirEntries, err := os.ReadDir(path)
	if err != nil {
		directory.Err = err
		return
	}
	lenDirEntries := len(dirEntries)

	localEventChan := make(chan model.Event)
	var localEventWG sync.WaitGroup

	entries := make([]model.Entry, 0, lenDirEntries)
	var totalSize int64
	for i := range lenDirEntries {
		info, err := dirEntries[i].Info()
		if err != nil {
			continue
		}

		name := info.Name()
		size := info.Size()
		mode := uint32(info.Mode())
		isDir := info.IsDir()

		totalSize += size

		if isDir {
			newDirectory := &model.EntryDirectory{
				Name:      name,
				TotalSize: size,
				Mode:      mode,
			}
			entries = append(entries, newDirectory)

			newDirectoryPath := filepath.Join(path, name)
			localEventWG.Add(1)
			go s.Read(newDirectoryPath, newDirectory, localEventChan, &localEventWG)

		} else {
			newFile := &model.EntryFile{
				Name: name,
				Size: size,
				Mode: mode,
			}
			entries = append(entries, newFile)
		}
	}
	lenEntries := len(entries)

	directory.Entries = entries
	directory.TotalEntryCount += lenEntries
	directory.TotalSize += totalSize

	eventUpdate := &model.EventUpdate{
		EntryCount: lenEntries,
		Size:       totalSize,
	}
	// go func() {
	eventChan <- eventUpdate
	// }()

	go util.CloseChan(localEventChan, &localEventWG)

	for localEventType := range localEventChan {
		switch localEvent := localEventType.(type) {

		case *model.EventUpdate:
			directory.TotalEntryCount += localEvent.EntryCount
			directory.TotalSize += localEvent.Size
		}

		// go func() {
		eventChan <- localEventType
		// }()
	}
}

func (s *service) GetMainDirectorySize(path string, directory *model.EntryDirectory) {
	info, err := os.Stat(path)
	if err != nil {
		directory.Err = err
		return
	}

	directory.TotalSize += info.Size()
}
