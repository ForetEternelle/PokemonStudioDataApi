package iter2

import (
	"slices"
	"testing"
)

func TestPeek(t *testing.T) {
	data := []int{1, 2, 3}
	it := slices.Values(data)

	var peeked []int
	peekedIt := Peek(it, func(v int) {
		peeked = append(peeked, v)
	})

	res := slices.Collect(peekedIt)
	if !slices.Equal(res, []int{1, 2, 3}) {
		t.Errorf("expected [1 2 3], got %v", res)
	}
	if !slices.Equal(peeked, []int{1, 2, 3}) {
		t.Errorf("expected peeked [1 2 3], got %v", peeked)
	}
}

func TestSkip(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	it := slices.Values(data)
	skipped := Skip(it, 3)
	res := slices.Collect(skipped)
	if !slices.Equal(res, []int{4, 5}) {
		t.Errorf("expected [4 5], got %v", res)
	}
}

func TestSkipAll(t *testing.T) {
	data := []int{1, 2, 3}
	it := slices.Values(data)
	skipped := Skip(it, 5)
	res := slices.Collect(skipped)
	if len(res) != 0 {
		t.Errorf("expected empty, got %v", res)
	}
}

func TestSkipZero(t *testing.T) {
	data := []int{1, 2, 3}
	it := slices.Values(data)
	skipped := Skip(it, 0)
	res := slices.Collect(skipped)
	if !slices.Equal(res, []int{1, 2, 3}) {
		t.Errorf("expected [1 2 3], got %v", res)
	}
}

func TestSkipUntil(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	it := slices.Values(data)
	skipped := SkipUntil(it, func(v int) bool { return v >= 3 })
	res := slices.Collect(skipped)
	if !slices.Equal(res, []int{3, 4, 5}) {
		t.Errorf("expected [3 4 5], got %v", res)
	}
}

func TestSkipUntilNeverMatch(t *testing.T) {
	data := []int{1, 2, 3}
	it := slices.Values(data)
	skipped := SkipUntil(it, func(v int) bool { return false })
	res := slices.Collect(skipped)
	if len(res) != 0 {
		t.Errorf("expected empty, got %v", res)
	}
}

func TestSkipUntilFirstMatch(t *testing.T) {
	data := []int{1, 2, 3}
	it := slices.Values(data)
	skipped := SkipUntil(it, func(v int) bool { return true })
	res := slices.Collect(skipped)
	if !slices.Equal(res, []int{1, 2, 3}) {
		t.Errorf("expected [1 2 3], got %v", res)
	}
}

func TestTake(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	it := slices.Values(data)
	res, hasMore := Take(it, 3)
	if !slices.Equal(res, []int{1, 2, 3}) {
		t.Errorf("expected [1 2 3], got %v", res)
	}
	if !hasMore {
		t.Error("expected hasMore to be true")
	}
}

func TestTakeNone(t *testing.T) {
	data := []int{1, 2, 3}
	it := slices.Values(data)
	res, hasMore := Take(it, 0)
	if len(res) != 0 {
		t.Errorf("expected empty, got %v", res)
	}
	if !hasMore {
		t.Error("expected hasMore to be true")
	}
}

func TestTakeMoreThanAvailable(t *testing.T) {
	data := []int{1, 2}
	it := slices.Values(data)
	res, hasMore := Take(it, 5)
	if !slices.Equal(res, []int{1, 2}) {
		t.Errorf("expected [1 2], got %v", res)
	}
	if hasMore {
		t.Error("expected hasMore to be false")
	}
}

func TestTakeExact(t *testing.T) {
	data := []int{1, 2, 3}
	it := slices.Values(data)
	res, hasMore := Take(it, 3)
	if !slices.Equal(res, []int{1, 2, 3}) {
		t.Errorf("expected [1 2 3], got %v", res)
	}
	if hasMore {
		t.Error("expected hasMore to be false")
	}
}
