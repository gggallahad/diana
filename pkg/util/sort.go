package util

import (
	"slices"

	"github.com/gggallahad/diana/internal/model"
)

func SortEntries(format model.SortFormat, entries []model.Entry) {
	switch format {

	case model.SortFormatNameAsc:
		sortByNameAsc(entries)
	case model.SortFormatNameDesc:
		sortByNameDesc(entries)
	case model.SortFormatSizeAsc:
		sortBySizeAsc(entries)
	case model.SortFormatSizeDesc:
		sortBySizeDesc(entries)
	}
}

func sortByNameAsc(entries []model.Entry) {
	slices.SortStableFunc(entries, func(a, b model.Entry) int {
		if a.GetName() < b.GetName() {
			return -1
		}
		if a.GetName() > b.GetName() {
			return 1
		}

		return 0
	})
}

func sortByNameDesc(entries []model.Entry) {
	slices.SortStableFunc(entries, func(a, b model.Entry) int {
		if a.GetName() > b.GetName() {
			return -1
		}
		if a.GetName() < b.GetName() {
			return 1
		}

		return 0
	})
}

func sortBySizeAsc(entries []model.Entry) {
	slices.SortStableFunc(entries, func(a, b model.Entry) int {
		if a.GetSize() < b.GetSize() {
			return -1
		}
		if a.GetSize() > b.GetSize() {
			return 1
		}

		return 0
	})
}

func sortBySizeDesc(entries []model.Entry) {
	slices.SortStableFunc(entries, func(a, b model.Entry) int {
		if a.GetSize() > b.GetSize() {
			return -1
		}
		if a.GetSize() < b.GetSize() {
			return 1
		}

		return 0
	})
}
