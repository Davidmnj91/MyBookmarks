package bookmarks

import (
	"github.com/Davidmnj91/MyBookmarks/domain/bookmarks"
	"github.com/gin-gonic/gin"
)

// AuthorValidator is a struct used to validate the JSON payload representing an author.
type BookmarkValidator struct {
	URL   string `binding:"required" json:"url"`
	Title string `json:"title"`
}

// Bind validates the JSON payload and returns a Author instance corresponding to the payload.
func Bind(c *gin.Context) (*bookmarks.Bookmark, error) {
	var json BookmarkValidator
	if err := c.ShouldBindJSON(&json); err != nil {
		return nil, err
	}

	bookmark := &bookmarks.Bookmark{
		Url:   json.URL,
		Title: json.Title,
	}

	return bookmark, nil
}
