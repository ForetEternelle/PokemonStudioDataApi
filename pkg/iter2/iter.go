package iter2

import "iter"

func Peek[V any](it iter.Seq[V], peek func(V)) iter.Seq[V] {
	return func(yield func(V) bool) {
		for item := range it {
			peek(item)
			if !yield(item) {
				break
			}
		}
	}
}

func Skip[V any](it iter.Seq[V], n int) iter.Seq[V] {
	return func(yield func(V) bool) {
		skipped := 0
		for item := range it {
			if skipped < n {
				skipped++
				continue
			}
			if !yield(item) {
				break
			}
		}
	}
}

func SkipUntil[V any](it iter.Seq[V], skip func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		skipping := true
		for item := range it {
			if skipping {
				if skip(item) {
					skipping = false
				}
				continue
			}
			if !yield(item) {
				break
			}
		}
	}
}

func Take[V any](it iter.Seq[V], n int) (result []V, hasMore bool) {
	count := 0
	for item := range it {
		if count >= n {
			return result, true
		}
		result = append(result, item)
		count++
	}

	return result, false
}
