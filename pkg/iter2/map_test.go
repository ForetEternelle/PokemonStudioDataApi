package iter2

import (
	"slices"
	"testing"
)

func TestMap(t *testing.T) {
	data := []int{1, 2, 3}
	it := slices.Values(data)
	mapped := Map(it, func(n int) string {
		return string(rune('a' + n - 1))
	})
	res := slices.Collect(mapped)
	if !slices.Equal(res, []string{"a", "b", "c"}) {
		t.Errorf("expected [a b c], got %v", res)
	}
}

func TestMapEmpty(t *testing.T) {
	it := slices.Values([]int{})
	mapped := Map(it, func(n int) int { return n * 2 })
	res := slices.Collect(mapped)
	if len(res) != 0 {
		t.Errorf("expected empty, got %v", res)
	}
}

func TestMap2(t *testing.T) {
	data := []int{1, 2, 3}
	it := slices.All(data)
	mapped := Map2(it, func(k, v int) (int, int) {
		return k, v * 10
	})
	i := 0
	expected := []struct{ k, v int }{{0, 10}, {1, 20}, {2, 30}}
	for k, v := range mapped {
		if k != expected[i].k || v != expected[i].v {
			t.Errorf("at %d: expected (%d, %d), got (%d, %d)", i, expected[i].k, expected[i].v, k, v)
		}
		i++
	}
	if i != 3 {
		t.Errorf("expected 3 items, got %d", i)
	}
}

func TestValues(t *testing.T) {
	data := []int{10, 20, 30}
	it := slices.All(data)
	vals := Values(it)
	res := slices.Collect(vals)
	if !slices.Equal(res, []int{10, 20, 30}) {
		t.Errorf("expected [10 20 30], got %v", res)
	}
}

func TestValuesEmpty(t *testing.T) {
	it := slices.All([]int{})
	vals := Values(it)
	res := slices.Collect(vals)
	if len(res) != 0 {
		t.Errorf("expected empty, got %v", res)
	}
}

func TestToSeq2(t *testing.T) {
	data := []int{1, 2, 3}
	it := slices.Values(data)
	seq2 := ToSeq2(it, func(n int) string {
		return string(rune('a' + n - 1))
	})
	i := 0
	expected := []struct{ k int; v string }{{1, "a"}, {2, "b"}, {3, "c"}}
	for k, v := range seq2 {
		if k != expected[i].k || v != expected[i].v {
			t.Errorf("at %d: expected (%d, %s), got (%d, %s)", i, expected[i].k, expected[i].v, k, v)
		}
		i++
	}
	if i != 3 {
		t.Errorf("expected 3 items, got %d", i)
	}
}
