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

func Or[T any](filters ...FilterFunc[T]) FilterFunc[T] {
	if len(filters) == 0 {
		return True[T]
	}
	return func(item T) bool {
		for _, filter := range filters {
			if filter(item) {
				return true
			}
		}
		return false
	}
}

func And[T any](filters ...FilterFunc[T]) FilterFunc[T] {
	if len(filters) == 0 {
		return True[T]
	}
	return func(item T) bool {
		for _, filter := range filters {
			if !filter(item) {
				return false
			}
		}
		return true
	}
}

func True[T any](T) bool {
	return true
}

func False[T any](T) bool {
	return false
}

func First[V any](it iter.Seq[V]) (V, bool) {
	for item := range it {
		return item, true
	}
	var zero V
	return zero, false
}

func Peek[V any](peek func(V), it iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for item := range it {
			peek(item)
			if !yield(item) {
				break
			}
		}
	}
}
