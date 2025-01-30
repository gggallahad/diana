package project

import (
	"sync"

	"github.com/gggallahad/diana/internal/model"
	"github.com/gggallahad/gui"
)

type (
	Handler interface {
		Init(ctx *gui.Context)

		Read(ctx *gui.Context)
		GetEvents(ctx *gui.Context)
		Draw(ctx *gui.Context)

		Kill(ctx *gui.Context, eventType gui.Event)
		Resize(ctx *gui.Context, eventType gui.Event)
		MoveCamera(ctx *gui.Context, eventType gui.Event)

		Move(ctx *gui.Context, eventType gui.Event)
	}

	Service interface {
		Read(path string, directory *model.EntryDirectory, eventChan chan<- model.Event, eventWG *sync.WaitGroup)
		GetMainDirectorySize(path string, directory *model.EntryDirectory)
	}
)
