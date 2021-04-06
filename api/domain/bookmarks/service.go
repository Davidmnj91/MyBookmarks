package bookmarks

type BookmarkService interface {
	CreateBookmark(*Bookmark) (*Bookmark, error)
	ReadBookmark(id string) (*Bookmark, error)
	EditBookmark(id string, bookmark *Bookmark) (*Bookmark, error)
	DeleteBookmark(id string) (*Bookmark, error)
	ListBookmarks() ([]Bookmark, error)
}

// Service struct handles author business logic tasks.
type Service struct {
	repository BookmarkRepository
}

func (svc *Service) CreateBookmark(bookmark *Bookmark) (*Bookmark, error) {
	return svc.repository.CreateBookmark(bookmark)
}

func (svc *Service) ReadBookmark(id string) (*Bookmark, error) {
	return svc.repository.ReadBookmark(id)
}

func (svc *Service) EditBookmark(id string, bookmark *Bookmark) (*Bookmark, error) {
	return svc.repository.EditBookmark(id, bookmark)
}

func (svc *Service) DeleteBookmark(id string) (*Bookmark, error) {
	return svc.repository.DeleteBookmark(id)
}

func (svc *Service) ListBookmarks() ([]Bookmark, error) {
	return svc.repository.ListBookmarks()
}

// NewService creates a new service struct
func NewService(repository BookmarkRepository) *Service {
	return &Service{repository: repository}
}
