package main

type (
	Entry interface {
		isEntry()
	}

	EntryDirectory struct {
		Name string

		Entries []Entry

		TotalEntryCount int
		TotalSize       int64
	}

	EntryFile struct {
		Name string
		Size int64
	}
)

func (e *EntryDirectory) isEntry() {
}

func (e *EntryFile) isEntry() {
}
