package bookmarks

import (
	domain "github.com/Davidmnj91/MyBookmarks/domain/bookmarks"
	domainErrors "github.com/Davidmnj91/MyBookmarks/domain/errors"
	"github.com/google/uuid"
)

func toDBModel(entity *domain.Bookmark) (*Bookmark, error) {
	id, err := uuid.Parse(entity.ID)

	if err != nil {
		appErr := domainErrors.NewAppErrorWithType(domainErrors.InvalidID)
		return nil, appErr
	}

	return &Bookmark{
		ID:    id,
		URL:   entity.Url,
		Title: entity.Title,
	}, nil
}

func toDomainModel(entity *Bookmark) *domain.Bookmark {
	return &domain.Bookmark{
		ID:        entity.ID.String(),
		Url:       entity.URL,
		Title:     entity.Title,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
