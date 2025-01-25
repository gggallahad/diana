package model

type (
	Event interface {
		isEvent()
	}

	EventUpdateEntryCount struct {
		EntryCount int
	}

	EventUpdateSize struct {
		Size int64
	}
)

func (e *EventUpdateEntryCount) isEvent() {
}

func (e *EventUpdateSize) isEvent() {
}
