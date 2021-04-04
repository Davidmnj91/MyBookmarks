package http

import (
	"github.com/Davidmnj91/MyBookmarks/domain/bookmarks"
	bookmarksRoutes "github.com/Davidmnj91/MyBookmarks/router/http/bookmarks"
	"github.com/Davidmnj91/MyBookmarks/router/http/errors"
	healthRoutes "github.com/Davidmnj91/MyBookmarks/router/http/health"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewHTTPHandler returns the HTTP requests handler
func NewHTTPHandler(bookmarkSvc bookmarks.BookmarkService) http.Handler {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	router.Use(errors.Handler)

	healthGroup := router.Group("/health")
	healthRoutes.NewRoutesFactory()(healthGroup)

	api := router.Group("/api")
	//api.Use(authMiddleware)

	bookmarksGroup := api.Group("/bookmarks")
	bookmarksRoutes.NewRoutesFactory(bookmarksGroup)(bookmarkSvc)

	return router
}
