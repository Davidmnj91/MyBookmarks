package bookmarks

import "time"

type Bookmark struct {
	ID    string
	Url   string
	Title string
	//TAGS      []string
	CreatedAt time.Time
	UpdatedAt time.Time
}
