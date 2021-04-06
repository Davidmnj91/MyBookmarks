package bookmarks

import (
	"github.com/Davidmnj91/MyBookmarks/domain/bookmarks"
	domainErrors "github.com/Davidmnj91/MyBookmarks/domain/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewRoutesFactory create and returns a factory to create routes for the authors
func NewRoutesFactory(group *gin.RouterGroup) func(service bookmarks.BookmarkService) {
	bookmarkRoutesFactory := func(service bookmarks.BookmarkService) {
		group.GET("/", func(c *gin.Context) {
			results, err := service.ListBookmarks()
			if err != nil {
				c.Error(err)
				return
			}

			var responseBookmarks = make([]BookmarkResponse, len(results))

			for i := range results {
				responseBookmarks[i] = *toResponseModel(&results[i])
			}

			response := &ListResponse{
				Data: responseBookmarks,
			}

			c.JSON(http.StatusOK, response)
		})

		group.POST("/", func(c *gin.Context) {
			bookmark, err := Bind(c)
			if err != nil {
				appError := domainErrors.NewAppError(err, domainErrors.ValidationError)
				c.Error(appError)
				return
			}

			newBookmark, err := service.CreateBookmark(bookmark)
			if err != nil {
				c.Error(err)
				return
			}

			c.JSON(http.StatusCreated, *toResponseModel(newBookmark))
		})

		group.PUT("/:bookmarkID", func(c *gin.Context) {
			id := c.Param("bookmarkID")

			bookmark, err := Bind(c)
			if err != nil {
				appError := domainErrors.NewAppError(err, domainErrors.ValidationError)
				c.Error(appError)
				return
			}

			editedBookmark, err := service.EditBookmark(id, bookmark)
			if err != nil {
				c.Error(err)
				return
			}

			c.JSON(http.StatusCreated, *toResponseModel(editedBookmark))
		})

		group.GET("/:bookmarkID", func(c *gin.Context) {
			id := c.Param("bookmarkID")
			result, err := service.ReadBookmark(id)
			if err != nil {
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, *toResponseModel(result))
		})

		group.DELETE("/:bookmarkID", func(c *gin.Context) {
			id := c.Param("bookmarkID")
			result, err := service.DeleteBookmark(id)
			if err != nil {
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, *toResponseModel(result))
		})
	}

	return bookmarkRoutesFactory
}
