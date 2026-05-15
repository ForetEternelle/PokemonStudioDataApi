package scroll

import (
	"iter"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
)

type Scroll[T any] struct {
	Content []T  `json:"content"`
	HasMore bool `json:"hasMore"`
}

func New[T any](content []T, hasMore bool) Scroll[T] {
	return Scroll[T]{
		Content: content,
		HasMore: hasMore,
	}
}

func Of[T any](it iter.Seq[T], size int) Scroll[T] {
	content, hasMore := iter2.Take(it, size)
	return New(content, hasMore)
}
