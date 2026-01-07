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

func TestCollect_FirstPage(t *testing.T) {
	content := []int{1, 2, 3, 4, 5}
	pr := NewPageRequest(0, 2)

	it := slices.Values(content)
	page := Collect(it, pr)

	expectedContent := []int{1, 2}
	if len(page.Content) != len(expectedContent) {
		t.Error("Page content len should be", len(expectedContent), ", has", len(page.Content))
	}

	for i, val := range expectedContent {
		if page.Content[i] != val {
			t.Error("Content at index", i, "should be", val, ", has", page.Content[i])
		}
	}

	if page.Total != len(content) {
		t.Error("Page total should be", len(content), ", has", page.Total)
	}
}

func TestCollect_LastPagePartial(t *testing.T) {
	content := []int{1, 2, 3, 4, 5}
	pr := NewPageRequest(2, 2)

	it := slices.Values(content)
	page := Collect(it, pr)

	expectedContent := []int{5}
	if len(page.Content) != len(expectedContent) {
		t.Error("Page content len should be", len(expectedContent), ", has", len(page.Content))
	}

	if page.Content[0] != expectedContent[0] {
		t.Error("Content should be", expectedContent[0], ", has", page.Content[0])
	}

	if page.Total != len(content) {
		t.Error("Page total should be", len(content), ", has", page.Total)
	}
}

func TestCollect_EmptyResult(t *testing.T) {
	content := []int{1, 2, 3}
	pr := NewPageRequest(2, 2)

	it := slices.Values(content)
	page := Collect(it, pr)

	if len(page.Content) != 0 {
		t.Error("Page content should be empty, has", len(page.Content))
	}

	if page.Total != len(content) {
		t.Error("Page total should be", len(content), ", has", page.Total)
	}
}

func TestCollect_EmptyIterator(t *testing.T) {
	var content []int
	pr := NewPageRequest(0, 2)

	it := slices.Values(content)
	page := Collect(it, pr)

	if len(page.Content) != 0 {
		t.Error("Page content should be empty, has", len(page.Content))
	}

	if page.Total != 0 {
		t.Error("Page total should be 0, has", page.Total)
	}
}

func TestNewPage(t *testing.T) {
	content := []string{"a", "b", "c"}
	page := NewPage(1, 10, content, 25)

	if page.Page != 1 {
		t.Error("Page should be 1, has", page.Page)
	}

	if page.Size != 10 {
		t.Error("Size should be 10, has", page.Size)
	}

	if len(page.Content) != len(content) {
		t.Error("Content length should be", len(content), ", has", len(page.Content))
	}

	if page.Total != 25 {
		t.Error("Total should be 25, has", page.Total)
	}

	for i, val := range content {
		if page.Content[i] != val {
			t.Error("Content at index", i, "should be", val, ", has", page.Content[i])
		}
	}
}

func TestNewPage_EmptyContent(t *testing.T) {
	var content []int
	page := NewPage(0, 5, content, 0)

	if page.Page != 0 {
		t.Error("Page should be 0, has", page.Page)
	}

	if page.Size != 5 {
		t.Error("Size should be 5, has", page.Size)
	}

	if len(page.Content) != 0 {
		t.Error("Content should be empty, has", len(page.Content))
	}

	if page.Total != 0 {
		t.Error("Total should be 0, has", page.Total)
	}
}

func TestHasNext(t *testing.T) {
	tests := []struct {
		name     string
		page     Page[int]
		expected bool
	}{
		{
			name: "has next - more items available",
			page: Page[int]{
				Page:    0,
				Size:    2,
				Content: []int{1, 2},
				Total:   5,
			},
			expected: true,
		},
		{
			name: "no next - last full page",
			page: Page[int]{
				Page:    1,
				Size:    2,
				Content: []int{3, 4},
				Total:   4,
			},
			expected: false,
		},
		{
			name: "no next - last partial page",
			page: Page[int]{
				Page:    2,
				Size:    2,
				Content: []int{5},
				Total:   5,
			},
			expected: false,
		},
		{
			name: "no next - empty total",
			page: Page[int]{
				Page:    0,
				Size:    2,
				Content: []int{},
				Total:   0,
			},
			expected: false,
		},
		{
			name: "has next - same start as total but content empty",
			page: Page[int]{
				Page:    2,
				Size:    2,
				Content: []int{},
				Total:   4,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.page.HasNext()
			if result != tt.expected {
				t.Error("HasNext should return", tt.expected, ", has", result)
			}
		})
	}
}

func TestHasPrevious(t *testing.T) {
	tests := []struct {
		name     string
		page     Page[int]
		expected bool
	}{
		{
			name: "has previous - page 1",
			page: Page[int]{
				Page:    1,
				Size:    2,
				Content: []int{3, 4},
				Total:   5,
			},
			expected: true,
		},
		{
			name: "no previous - first page",
			page: Page[int]{
				Page:    0,
				Size:    2,
				Content: []int{1, 2},
				Total:   5,
			},
			expected: false,
		},
		{
			name: "has previous - deep page",
			page: Page[int]{
				Page:    5,
				Size:    10,
				Content: []int{51, 52, 53},
				Total:   100,
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.page.HasPrevious()
			if result != tt.expected {
				t.Error("HasPrevious should return", tt.expected, ", has", result)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		page     Page[int]
		expected bool
	}{
		{
			name: "empty - zero total",
			page: Page[int]{
				Page:    0,
				Size:    2,
				Content: []int{},
				Total:   0,
			},
			expected: true,
		},
		{
			name: "not empty - has total",
			page: Page[int]{
				Page:    0,
				Size:    2,
				Content: []int{1, 2},
				Total:   5,
			},
			expected: false,
		},
		{
			name: "not empty - content empty but has total",
			page: Page[int]{
				Page:    2,
				Size:    2,
				Content: []int{},
				Total:   3,
			},
			expected: false,
		},
		{
			name: "empty - no content and zero total",
			page: Page[int]{
				Page:    1,
				Size:    5,
				Content: []int{},
				Total:   0,
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.page.IsEmpty()
			if result != tt.expected {
				t.Error("IsEmpty should return", tt.expected, ", has", result)
			}
		})
	}
}
