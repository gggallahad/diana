package util

import (
	"github.com/gggallahad/diana/internal/model"
	"sync"
)

func CloseChan(eventChan chan model.Event, eventWG *sync.WaitGroup) {
	eventWG.Wait()
	close(eventChan)
}
