package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginationParams struct {
	CurrentPage int
	PerPage     int
}

func GetPaginationParams(c *gin.Context) PaginationParams {
	currentPage := 1
	perPage := 10

	if cp, err := strconv.Atoi(c.Query("currentPage")); err == nil {
		currentPage = cp
	}
	if pp, err := strconv.Atoi(c.Query("perPage")); err == nil {
		perPage = pp
	}

	return PaginationParams{
		CurrentPage: currentPage,
		PerPage:     perPage,
	}
}
