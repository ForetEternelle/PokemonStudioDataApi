package pagination

import (
	"slices"
	"testing"
)

func TestCollect(t *testing.T) {
	content := []int{1, 2, 3, 4}
	contentLen := len(content)
	pr := NewPageRequest(1, 2)

	it := slices.Values(content)
	page := Collect(it, pr)

	resLen := len(page.Content)
	expectedLen := 2

	if resLen != expectedLen {
		t.Error("Page content len should be", expectedLen, ", has", resLen)
	}

	if page.Total != contentLen {
		t.Error("Page total should be", contentLen, ", has", page.Total)
	}

	if page.Page != pr.Page {
		t.Error("Page should be", pr.Page, ", has", page.Page)
	}

	if page.Size != pr.Size {
		t.Error("Page size should be", pr.Size, ", has", page.Size)
	}

}