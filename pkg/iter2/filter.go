package iter2

import "iter"

type FilterFunc[T any] func(T) bool

func Filter[T any](filter FilterFunc[T], it iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range it {
			if filter(item) {
				if !yield(item) {
					break
				}
			}
		}
	}
}
