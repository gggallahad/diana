package model

type (
	Entry interface {
		isEntry()
		// Done()

		GetName() string
		GetSize() int64
	}

	EntryDirectory struct {
		Name string

		Entries         []Entry
		TotalEntryCount int
		TotalSize       int64
		Mode            uint32

		IsDone bool

		Err error
	}

	EntryFile struct {
		Name string
		Size int64
		Mode uint32

		// IsDone bool

		// Err error
	}
)

func (e *EntryDirectory) Done() {
	e.IsDone = true
}

func (e *EntryDirectory) GetName() string {
	return e.Name
}

func (e *EntryFile) GetName() string {
	return e.Name
}

func (e *EntryDirectory) GetSize() int64 {
	return e.TotalSize
}

func (e *EntryFile) GetSize() int64 {
	return e.Size
}

func (e *EntryDirectory) isEntry() {
}

func (e *EntryFile) isEntry() {
}
