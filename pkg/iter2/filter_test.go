package iter2

import (
	"slices"
	"testing"
)

func TestFilter(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	it := slices.Values(data)

	moreThan3 := func(n int) bool {
		return n > 3
	}

	res := slices.Collect(Filter(moreThan3, it))
	if len(res) != 2 {
		t.Error("Res lenght should be 2")
	}

	if res[0] != 4 {
		t.Error("Res should contains 4")
	}
	if res[1] != 5 {
		t.Error("Res should contains 5")
	}
}

func TestOr(t *testing.T) {
	even := func(n int) bool { return n%2 == 0 }
	greaterThan4 := func(n int) bool { return n > 4 }
	or := Or(even, greaterThan4)

	data := []int{1, 2, 3, 4, 5, 6}
	it := slices.Values(data)
	res := slices.Collect(Filter(or, it))

	if len(res) != 4 {
		t.Errorf("expected 4, got %d", len(res))
	}
}

func TestOrEmpty(t *testing.T) {
	or := Or[int]()
	data := []int{1, 2, 3}
	it := slices.Values(data)
	res := slices.Collect(Filter(or, it))

	if len(res) != 3 {
		t.Errorf("expected 3, got %d", len(res))
	}
}

func TestAnd(t *testing.T) {
	even := func(n int) bool { return n%2 == 0 }
	greaterThan2 := func(n int) bool { return n > 2 }
	and := And(even, greaterThan2)

	data := []int{1, 2, 3, 4, 5, 6}
	it := slices.Values(data)
	res := slices.Collect(Filter(and, it))

	if len(res) != 2 {
		t.Errorf("expected 2, got %d", len(res))
	}
	if res[0] != 4 || res[1] != 6 {
		t.Error("expected [4, 6]")
	}
}

func TestAndEmpty(t *testing.T) {
	and := And[int]()
	data := []int{1, 2, 3}
	it := slices.Values(data)
	res := slices.Collect(Filter(and, it))

	if len(res) != 3 {
		t.Errorf("expected 3, got %d", len(res))
	}
}

func TestTrue(t *testing.T) {
	if !True(42) {
		t.Error("True should always return true")
	}
}

func TestFalse(t *testing.T) {
	if False(42) {
		t.Error("False should always return false")
	}
}
