package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	path string
)

func init() {
	flag.StringVar(&path, "path", ".", "path to directory")

	flag.Parse()
}

func main() {
	start := time.Now()

	mainEntry := EntryDirectory{
		Name: path,
	}

	eventChan := make(chan Event)
	var eventWG sync.WaitGroup

	eventWG.Add(1)
	go closeChan(eventChan, &eventWG)

	go Read(&mainEntry, eventChan, &eventWG)

	for eventType := range eventChan {
		switch event := eventType.(type) {

		case EventError:
			log.Println(event.Err)

		case EventUpdateEntryCount:

		case EventUpdateSize:

		}
	}

	end := time.Since(start)

	log.Printf("Directory name: %s   Total size: %d   Total entry count: %d", mainEntry.Name, mainEntry.TotalSize, mainEntry.TotalEntryCount)
	log.Printf("Time: %.3f seconds", end.Seconds())
}

func Read(entryType Entry, eventChan chan<- Event, eventWG *sync.WaitGroup) {
	defer eventWG.Done()

	switch entry := entryType.(type) {

	case *EntryFile:
		info, err := os.Stat(entry.Name)
		if err != nil {
			eventError := EventError{
				Err: errors.Join(ErrGetFileInfo, err),
			}
			eventChan <- eventError
			return
		}

		entry.Size = info.Size()

		eventUpdateSize := EventUpdateSize{
			Size: info.Size(),
		}
		eventChan <- eventUpdateSize

	case *EntryDirectory:
		dirEntries, err := os.ReadDir(entry.Name)
		if err != nil {
			eventError := EventError{
				Err: errors.Join(ErrReadDirectory, err),
			}
			eventChan <- eventError
			return
		}
		lenDirEntries := len(dirEntries)

		entry.Entries = make([]Entry, lenDirEntries)
		entry.TotalEntryCount += lenDirEntries

		eventUpdateElementCount := EventUpdateEntryCount{
			EntryCount: lenDirEntries,
		}
		eventChan <- eventUpdateElementCount

		localEventChan := make(chan Event)
		var localEventWG sync.WaitGroup

		localEventWG.Add(lenDirEntries)
		go closeChan(localEventChan, &localEventWG)

		for i := range lenDirEntries {
			entryPath := filepath.Join(entry.Name, dirEntries[i].Name())
			if !dirEntries[i].IsDir() {
				entry.Entries[i] = &EntryFile{
					Name: entryPath,
				}
			} else {
				entry.Entries[i] = &EntryDirectory{
					Name: entryPath,
				}
			}

			go Read(entry.Entries[i], localEventChan, &localEventWG)
		}

		for localEventType := range localEventChan {
			switch localEvent := localEventType.(type) {
			case EventError:
			case EventUpdateEntryCount:
				entry.TotalEntryCount += localEvent.EntryCount
			case EventUpdateSize:
				entry.TotalSize += localEvent.Size
			}

			eventChan <- localEventType
		}
	}
}
