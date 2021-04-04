package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"

	domain "github.com/Davidmnj91/MyBookmarks/domain/errors"
)

// Handler is Gin middleware to handle errors.
func Handler(c *gin.Context) {
	// Execute request handlers and then handle any errors
	c.Next()

	errs := c.Errors

	if len(errs) > 0 {
		err, ok := errs[0].Err.(*domain.AppError)
		if ok {
			switch err.Type {
			case domain.InvalidID:
				c.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		// Error is not AppError, return a generic internal server error
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
}
