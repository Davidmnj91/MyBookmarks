package bookmarks

import domain "github.com/Davidmnj91/MyBookmarks/domain/bookmarks"

func toResponseModel(entity *domain.Bookmark) *BookmarkResponse {
	return &BookmarkResponse{
		ID:        entity.ID,
		URL:       entity.Url,
		Title:     entity.Title,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
