package model

type (
	Event interface {
		isEvent()
	}

	EventUpdate struct {
		EntryCount int
		Size       int64
	}

	// EventUpdateSize struct {
	// }
)

func (e *EventUpdate) isEvent() {
}

// func (e *EventUpdateSize) isEvent() {
// }
