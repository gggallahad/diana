package main

import (
	"flag"
	"log"
	"path/filepath"
	"runtime"

	"github.com/gggallahad/diana/internal/handler"
	"github.com/gggallahad/diana/internal/model"
	"github.com/gggallahad/diana/internal/service"
	"github.com/gggallahad/diana/pkg/util"
	"github.com/gggallahad/gui"
)

func init() {
	flag.StringVar(&util.Path, util.PathFlagName, util.DefaultPath, util.PathDescription)
	flag.Int64Var(&util.FPS, util.FPSFlagName, util.DefaultFPS, util.FPSDescription)
	flag.IntVar(&util.NumCPU, util.NumCPUFlagName, util.DefaultNumCPU, util.NumCPUDescription)
	flag.StringVar(&util.SortFormatString, util.SortFormatFlagName, util.DefaultSortFormat, util.SortFormatDescription)
	flag.StringVar(&util.SizeFormatString, util.SizeFormatFlagName, util.DefaultSizeFormat, util.SizeFormatDescription)

	flag.Parse()

	if util.FPS < 0 {
		util.FPS = 0
	}

	maxNumCPU := runtime.NumCPU()
	if util.NumCPU > maxNumCPU {
		util.NumCPU = maxNumCPU
	}

	switch util.SortFormatString {
	case "aname":
		util.SortFormat = model.SortFormatNameAsc
	case "dname":
		util.SortFormat = model.SortFormatNameDesc
	case "asize":
		util.SortFormat = model.SortFormatSizeAsc
	case "dsize":
		util.SortFormat = model.SortFormatSizeDesc
	default:
		util.SortFormat = model.SortFormatSizeDesc
	}

	switch util.SizeFormatString {
	case "dynamic":
		util.SizeFormat = model.SizeFormatDynamic
	case "b":
		util.SizeFormat = model.SizeFormatBytes
	case "kb":
		util.SizeFormat = model.SizeFormatKilobytes
	case "mb":
		util.SizeFormat = model.SizeFormatMegabytes
	case "gb":
		util.SizeFormat = model.SizeFormatGigabytes
	case "tb":
		util.SizeFormat = model.SizeFormatTerabytes
	default:
		util.SizeFormat = model.SizeFormatDynamic
	}
}

func main() {
	runtime.GOMAXPROCS(util.NumCPU)

	mainDirectoryName := filepath.Base(util.Path)

	newService := service.NewService()
	newHandler := handler.NewHandler(mainDirectoryName, newService)

	screen, err := gui.NewScreen()
	if err != nil {
		log.Println(err)
		return
	}

	err = screen.Init()
	if err != nil {
		log.Println(err)
		return
	}
	defer screen.Close()

	screen.BindInitHandlers(newHandler.Init)

	screen.BindGlobalMiddlewares(newHandler.Kill, newHandler.Resize, newHandler.MoveCamera)

	screen.BindHandlers(gui.NoState, newHandler.Move)

	screen.BindGlobalPostwares(newHandler.DrawInterface)

	screen.BindBackgroundHandlers(newHandler.Read, newHandler.GetEvents, newHandler.Draw)

	screen.Run()
}

// func main() {
// 	start := time.Now()

// 	mainEntry := &model.EntryDirectory{
// 		Name: path,
// 	}

// 	eventChan := make(chan model.Event)
// 	var eventWG sync.WaitGroup

// 	eventWG.Add(1)
// 	go util.CloseChan(eventChan, &eventWG)

// 	go Read(mainEntry, eventChan, &eventWG)

// 	for eventType := range eventChan {
// 		switch event := eventType.(type) {

// 		case *model.EventError:
// 			log.Println(event.Err)

// 		case *model.EventUpdateEntryCount:

// 		case *model.EventUpdateSize:

// 		}
// 	}

// 	end := time.Since(start)

// 	log.Printf("Directory name: %s   Total size: %d   Total entry count: %d", mainEntry.Name, mainEntry.TotalSize, mainEntry.TotalEntryCount)
// 	log.Printf("Time: %.3f seconds", end.Seconds())
// }

// func Read(entryType model.Entry, eventChan chan<- model.Event, eventWG *sync.WaitGroup) {
// 	defer eventWG.Done()

// 	switch entry := entryType.(type) {

// 	case *model.EntryFile:
// 		info, err := os.Stat(entry.Name)
// 		if err != nil {
// 			if errors.Is(err, os.ErrNotExist) {
// 				return
// 			}
// 			eventError := &model.EventError{
// 				Err: errors.Join(util.ErrGetFileInfo, err),
// 			}
// 			eventChan <- eventError
// 			return
// 		}

// 		entry.Size = info.Size()

// 		eventUpdateSize := &model.EventUpdateSize{
// 			Size: info.Size(),
// 		}
// 		eventChan <- eventUpdateSize

// 	case *model.EntryDirectory:
// 		dirEntries, err := os.ReadDir(entry.Name)
// 		if err != nil {
// 			if errors.Is(err, os.ErrNotExist) {
// 				return
// 			}
// 			eventError := &model.EventError{
// 				Err: errors.Join(util.ErrReadDirectory, err),
// 			}
// 			eventChan <- eventError
// 			return
// 		}
// 		lenDirEntries := len(dirEntries)

// 		entry.Entries = make([]model.Entry, lenDirEntries)
// 		entry.TotalEntryCount += lenDirEntries

// 		eventUpdateElementCount := &model.EventUpdateEntryCount{
// 			EntryCount: lenDirEntries,
// 		}
// 		eventChan <- eventUpdateElementCount

// 		localEventChan := make(chan model.Event)
// 		var localEventWG sync.WaitGroup

// 		localEventWG.Add(lenDirEntries)
// 		go util.CloseChan(localEventChan, &localEventWG)

// 		entries := entry.Entries
// 		for i := range lenDirEntries {
// 			_ = entries[i] //bc

// 			entryPath := filepath.Join(entry.Name, dirEntries[i].Name())
// 			if !dirEntries[i].IsDir() {
// 				entries[i] = &model.EntryFile{
// 					Name: entryPath,
// 				}
// 			} else {
// 				entries[i] = &model.EntryDirectory{
// 					Name: entryPath,
// 				}
// 			}

// 			go Read(entries[i], localEventChan, &localEventWG)
// 		}

// 		for localEventType := range localEventChan {
// 			switch localEvent := localEventType.(type) {
// 			case *model.EventError:
// 			case *model.EventUpdateEntryCount:
// 				entry.TotalEntryCount += localEvent.EntryCount
// 			case *model.EventUpdateSize:
// 				entry.TotalSize += localEvent.Size
// 			}

// 			eventChan <- localEventType
// 		}
// 	}
// }
