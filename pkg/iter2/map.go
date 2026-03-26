package iter2

import "iter"

type MapFunc[T any, R any] func(T) R
type MapFunc2[K any, T any, R any] func(K, T) R

func Map[T any, R any](fn MapFunc[T, R], it iter.Seq[T]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for item := range it {
			if !yield(fn(item)) {
				break
			}
		}
	}
}

func Map2[K any, V any, R any](fn MapFunc2[K, V, R], it iter.Seq2[K, V]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for k, value := range it {
			if !yield(fn(k, value)) {
				break
			}
		}
	}
}
