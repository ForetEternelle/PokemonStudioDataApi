package pagination

import (
	"iter"
)

type Page[T any] struct {
	Page    int `json:"page"`
	Size    int `json:"size"`
	Content []T `json:"content"`
	Total   int `json:"total"`
}

func NewPage[T any](page int, size int, content []T, total int) Page[T] {

	return Page[T]{
		Page:    page,
		Size:    size,
		Content: content,
		Total:   total,
	}
}

func Collect[T any](it iter.Seq[T], pageRequest PageRequest) Page[T] {
	var content = make([]T, 0)
	offset := pageRequest.Size * pageRequest.Page
	total := 0
	skip := 0
	found := 0

	for item := range it {
		total++
		if skip < offset {
			skip++
			continue
		}

		if found < pageRequest.Size {
			copy := item
			content = append(content, copy)
			found++
		}
	}
	return NewPage(pageRequest.Page, pageRequest.Size, content, total)
}
