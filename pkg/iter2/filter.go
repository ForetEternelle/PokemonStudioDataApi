package iter2

import "iter"

// FilterFunc is a predicate function used to test items of type T.
type FilterFunc[T any] func(T) bool

// Filter returns a Seq containing only items from it for which filter returns true.
func Filter[T any](it iter.Seq[T], filter FilterFunc[T]) iter.Seq[T] {
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

// Or returns a FilterFunc that returns true if any of the provided filters returns true.
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

// And returns a FilterFunc that returns true only if all of the provided filters return true.
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

// True returns a FilterFunc that always returns true regardless of the input.
func True[T any](T) bool {
	return true
}

// False returns a FilterFunc that always returns false regardless of the input.
func False[T any](T) bool {
	return false
}

// First returns the first item from the iterator and true, or the zero value and false if the iterator is empty.
func First[V any](it iter.Seq[V]) (V, bool) {
	for item := range it {
		return item, true
	}
	var zero V
	return zero, false
}
