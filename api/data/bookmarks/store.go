package bookmarks

import (
	errors2 "errors"
	domain "github.com/Davidmnj91/MyBookmarks/domain/bookmarks"
	domainErrors "github.com/Davidmnj91/MyBookmarks/domain/errors"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const (
	createError = "Error in creating new bookmark"
	editError   = "Error in editing bookmark"
	deleteError = "Error in deleting bookmark"
	readError   = "Error in finding bookmark in the database"
	listError   = "Error in getting bookmarks from the database"
)

type Store struct {
	db *gorm.DB
}

func New(db *gorm.DB) (*Store, error) {
	if err := db.AutoMigrate(&Bookmark{}); err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}

func (s *Store) CreateBookmark(bookmark *domain.Bookmark) (*domain.Bookmark, error) {
	entity, err := toDBModel(bookmark)

	if err != nil {
		return nil, err
	}

	if err := s.db.Create(entity).Error; err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, createError), domainErrors.RepositoryError)
		return nil, appErr
	}

	return toDomainModel(entity), nil
}

func (s *Store) ReadBookmark(id string) (*domain.Bookmark, error) {
	result := &Bookmark{}

	if query := s.db.Where("id = ?", id).First(result); query.Error != nil {
		if errors2.Is(query.Error, gorm.ErrRecordNotFound) {
			appErr := domainErrors.NewAppErrorWithType(domainErrors.NotFound)
			return nil, appErr
		}

		appErr := domainErrors.NewAppError(errors.Wrap(query.Error, readError), domainErrors.RepositoryError)
		return nil, appErr
	}

	return toDomainModel(result), nil
}

func (s *Store) EditBookmark(id string, bookmark *domain.Bookmark) (*domain.Bookmark, error) {
	bookmark.ID = id
	entity, err := toDBModel(bookmark)

	if err != nil {
		return nil, err
	}

	if err := s.db.Where("id = ?", id).Updates(entity).Error; err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, editError), domainErrors.RepositoryError)
		return nil, appErr
	}

	return toDomainModel(entity), nil
}

func (s *Store) DeleteBookmark(id string) (*domain.Bookmark, error) {
	result := &Bookmark{}

	if err := s.db.Where("id = ?", id).Delete(result).Error; err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, deleteError), domainErrors.RepositoryError)
		return nil, appErr
	}

	return toDomainModel(result), nil
}

func (s *Store) ListBookmarks() ([]domain.Bookmark, error) {
	var results []Bookmark

	if err := s.db.Find(&results).Error; err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, listError), domainErrors.RepositoryError)
		return nil, appErr
	}

	var bookmarks = make([]domain.Bookmark, len(results))

	for i := range results {
		bookmarks[i] = *toDomainModel(&results[i])
	}

	return bookmarks, nil
}
