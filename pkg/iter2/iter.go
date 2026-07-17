package iter2

import "iter"

// Peek returns a Seq that calls peek for each item before yielding it, allowing side effects without consuming the item.
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

// Skip returns a Seq that skips the first n items from it before yielding the rest.
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

// SkipUntil returns a Seq that skips items until the skip predicate returns true, then yields all remaining items.
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

// Take collects up to n items from it into a slice. It returns the collected items and whether the iterator had more items remaining.
func Take[V any](it iter.Seq[V], n int) (result []V, hasMore bool) {
	result = []V{}
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
