package pagination

type PageRequest struct {
	Page int `json:"page" minimum:"0"`
	Size int `json:"size" minimum:"-1"`
}

// NewPageRequest create a page request of the given size
// page the number of the page, starting at 0
// size the size of the page
func NewPageRequest(page int, size int) PageRequest {
	return PageRequest{
		Page: page,
		Size: size,
	}
}
