package iter2

import "iter"

type MapFunc[T any, R any] func(T) R
type MapFunc2[T any, R any] func(int, T) R

func Map[T any, R any](fn MapFunc[T, R], it iter.Seq[T]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for item := range it {
			if !yield(fn(item)) {
				break
			}
		}
	}
}

func Map2[T any, R any](fn MapFunc2[T, R], it iter.Seq[T]) iter.Seq[R] {
	i := 0
	return func(yield func(R) bool) {
		for item := range it {
			if !yield(fn(i, item)) {
				break
			}
			i++
		}
	}
}
