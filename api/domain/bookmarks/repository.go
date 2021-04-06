package bookmarks

type BookmarkRepository interface {
	// BookmarkRepository provides an abstraction on top of the bookmark data source
	CreateBookmark(*Bookmark) (*Bookmark, error)
	ReadBookmark(string) (*Bookmark, error)
	EditBookmark(string, *Bookmark) (*Bookmark, error)
	DeleteBookmark(string) (*Bookmark, error)
	ListBookmarks() ([]Bookmark, error)
}
