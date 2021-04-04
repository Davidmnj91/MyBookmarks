package bookmarks

import "time"

// AuthorResponse struct defines response fields
type BookmarkResponse struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	Title     string    `json:"Title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ListResponse struct defines authors list response structure
type ListResponse struct {
	Data []BookmarkResponse `json:"data"`
}
