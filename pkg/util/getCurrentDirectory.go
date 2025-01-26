package util

// func GetCurrentDirectory(mainDirectory *model.EntryDirectory, pastDirectories []int) (*model.EntryDirectory, error) {
// 	currentDirectory := mainDirectory

// 	for i := range pastDirectories {
// 		if len(currentDirectory.Entries)-1 < pastDirectories[i] {
// 			return nil, ErrNoEntry
// 		}

// 		entry := currentDirectory.Entries[pastDirectories[i]]
// 		nextDirectory, ok := entry.(*model.EntryDirectory)
// 		if !ok {
// 			return nil, ErrEntryNotADirectory
// 		}

// 		currentDirectory = nextDirectory
// 	}

// 	return currentDirectory, nil
// }
