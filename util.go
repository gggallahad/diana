package main

import "sync"

func closeChan(eventChan chan Event, eventWG *sync.WaitGroup) {
	eventWG.Wait()
	close(eventChan)
}
