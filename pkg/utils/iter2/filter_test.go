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