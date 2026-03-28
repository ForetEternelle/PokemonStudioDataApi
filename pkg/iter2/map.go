package iter2

import "iter"

type MapFunc[T any, R any] func(T) R
type MapFunc2[K, V, X, Y any] func(K, V) (X, Y)

// Map applies a mapping function to each item in a Seq and returns a Seq of the results.
func Map[T any, R any](fn MapFunc[T, R], it iter.Seq[T]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for item := range it {
			if !yield(fn(item)) {
				break
			}
		}
	}
}

// Map2 applies a mapping function to each key-value pair in a Seq2 and returns a Seq2 of the results.
func Map2[K, V, X, Y any](fn MapFunc2[K, V, X, Y], it iter.Seq2[K, V]) iter.Seq2[X, Y] {
	return func(yield func(X, Y) bool) {
		for k, value := range it {
			if !yield(fn(k, value)) {
				break
			}
		}
	}
}

// Values extracts the values from a Seq2 and returns them as a Seq.
func Values[K, V any](it iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, value := range it {
			if !yield(value) {
				break
			}
		}
	}
}

// ToSeq2 converts a Seq of values into a Seq2 of values and their mapped results using the provided mapping function.
func ToSeq2[V, R any, I iter.Seq[V]](it I, mapper MapFunc[V, R]) iter.Seq2[V, R] {
	return func(yield func(V, R) bool) {
		for item := range it {
			if !yield(item, mapper(item)) {
				break
			}
		}
	}
}
