package service

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/gggallahad/diana/internal/model"
	"github.com/gggallahad/diana/pkg/util"
)

func (s *service) Read(path string, entryType model.Entry, eventChan chan<- model.Event, eventWG *sync.WaitGroup) {
	defer eventWG.Done()
	defer entryType.Done()

	switch entry := entryType.(type) {

	case *model.EntryFile:
		info, err := os.Stat(path)
		if err != nil {
			entry.Err = err
			return
		}

		entry.Size = info.Size()

		eventUpdateSize := &model.EventUpdateSize{
			Size: info.Size(),
		}
		eventChan <- eventUpdateSize

	case *model.EntryDirectory:
		dirEntries, err := os.ReadDir(path)
		if err != nil {
			entry.Err = err
			return
		}
		lenDirEntries := len(dirEntries)

		entry.Entries = make([]model.Entry, lenDirEntries)
		entry.TotalEntryCount += lenDirEntries

		eventUpdateElementCount := &model.EventUpdateEntryCount{
			EntryCount: lenDirEntries,
		}
		eventChan <- eventUpdateElementCount

		localEventChan := make(chan model.Event)
		var localEventWG sync.WaitGroup

		localEventWG.Add(lenDirEntries)
		go util.CloseChan(localEventChan, &localEventWG)

		entries := entry.Entries
		for i := range lenDirEntries {
			_ = entries[i] //bc

			entryName := dirEntries[i].Name()
			entryPath := filepath.Join(path, entryName)
			if !dirEntries[i].IsDir() {
				entries[i] = &model.EntryFile{
					Name: entryName,
				}
			} else {
				entries[i] = &model.EntryDirectory{
					Name: entryName,
				}
			}

			go s.Read(entryPath, entries[i], localEventChan, &localEventWG)
		}

		for localEventType := range localEventChan {
			switch localEvent := localEventType.(type) {
			case *model.EventUpdateEntryCount:
				entry.TotalEntryCount += localEvent.EntryCount
			case *model.EventUpdateSize:
				entry.TotalSize += localEvent.Size
			}

			eventChan <- localEventType
		}
	}
}
