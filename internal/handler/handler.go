package handler

import (
	"sync"
	"time"

	"github.com/gggallahad/diana/internal/model"
	"github.com/gggallahad/diana/internal/project"
	"github.com/gggallahad/diana/pkg/util"
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

		currentCursorPositionY int
		previousDirectories    []*model.EntryDirectory
		// isCurrentDirectoryUpToDate bool

		sizeFormat model.SizeFormat
		sortFormat model.SortFormat

		viewPositionX int
		viewPositionY int

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

	currentCursorPositionY := 0
	previousDirectories := []*model.EntryDirectory{mainDirectory}
	// isCurrentDirectoryUpToDate := false

	sizeFormat := util.SizeFormat
	sortFormat := util.SortFormat

	viewPositionX := 0
	viewPositionY := 0

	return &handler{
		mainDirectory:          mainDirectory,
		eventChan:              eventChan,
		stopDraw:               stopDraw,
		currentCursorPositionY: currentCursorPositionY,
		previousDirectories:    previousDirectories,
		// isCurrentDirectoryUpToDate: isCurrentDirectoryUpToDate,
		sizeFormat:    sizeFormat,
		sortFormat:    sortFormat,
		viewPositionX: viewPositionX,
		viewPositionY: viewPositionY,
		Service:       service,
	}
}
