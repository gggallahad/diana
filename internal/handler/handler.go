package handler

import (
	"sync"
	"time"

	"github.com/gggallahad/diana/internal/model"
	"github.com/gggallahad/diana/internal/project"
)

type (
	handler struct {
		mainDirectory *model.EntryDirectory

		eventChan chan model.Event
		eventWG   sync.WaitGroup

		drawTick <-chan time.Time
		stopDraw chan struct{}

		startTime time.Time
		endTime   time.Time

		// cursorPosition   int
		pastDirectories            []int
		currentDirectory           *model.EntryDirectory
		isCurrentDirectoryUpToDate bool

		sizeFormat model.SizeFormat

		Service project.Service

		drawMutex sync.Mutex
	}
)

func NewHandler(mainDirectoryName string, service project.Service) project.Handler {
	mainDirectory := &model.EntryDirectory{
		Name: mainDirectoryName,
	}

	eventChan := make(chan model.Event)

	stopDraw := make(chan struct{})

	pastDirectories := make([]int, 0)

	currentDirectory := mainDirectory

	isCurrentDirectoryUpToDate := false

	sizeFormat := model.SizeFormatDynamic

	return &handler{
		mainDirectory:              mainDirectory,
		eventChan:                  eventChan,
		stopDraw:                   stopDraw,
		pastDirectories:            pastDirectories,
		currentDirectory:           currentDirectory,
		isCurrentDirectoryUpToDate: isCurrentDirectoryUpToDate,
		sizeFormat:                 sizeFormat,
		Service:                    service,
	}
}
