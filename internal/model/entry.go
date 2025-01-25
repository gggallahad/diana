package model

type (
	Entry interface {
		isEntry()
		Done()
	}

	EntryDirectory struct {
		Name string

		Entries         []Entry
		TotalEntryCount int
		TotalSize       int64

		IsDone bool

		Err error
	}

	EntryFile struct {
		Name string
		Size int64

		IsDone bool

		Err error
	}
)

func (e *EntryDirectory) isEntry() {
}

func (e *EntryFile) isEntry() {
}

func (e *EntryDirectory) Done() {
	e.IsDone = true
}

func (e *EntryFile) Done() {
	e.IsDone = true
}
