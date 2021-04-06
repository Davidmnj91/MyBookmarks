package main

import (
	"net/http"
	"strconv"

	"github.com/Davidmnj91/MyBookmarks/config"
	"github.com/Davidmnj91/MyBookmarks/data/database"

	"github.com/Davidmnj91/MyBookmarks/domain/bookmarks"

	router "github.com/Davidmnj91/MyBookmarks/router/http"

	bookmarksStore "github.com/Davidmnj91/MyBookmarks/data/bookmarks"
)

func main() {
	configuration := config.NewConfig()

	// establish DB connection
	db, err := database.Connect(configuration.Database)
	if err != nil {
		panic(err)
	}

	bookmarksRepo, err := bookmarksStore.New(db)

	if err != nil {
		panic(err)
	}

	bookmarksSvc := bookmarks.NewService(bookmarksRepo)

	httpRouter := router.NewHTTPHandler(bookmarksSvc)
	err = http.ListenAndServe(":"+strconv.Itoa(configuration.Port), httpRouter)
	if err != nil {
		panic(err)
	}
}
