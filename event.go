package main

type (
	Event interface {
		isEvent()
	}

	EventError struct {
		Err error
	}

	EventUpdateEntryCount struct {
		EntryCount int
	}

	EventUpdateSize struct {
		Size int64
	}
)

func (e EventError) isEvent() {
}

func (e EventUpdateEntryCount) isEvent() {
}

func (e EventUpdateSize) isEvent() {
}
