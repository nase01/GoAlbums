package helpers

import (
	"time"

	"github.com/gin-gonic/gin"
)

type QueryFilters struct {
	Pagination PaginationParams
	Sort       string
	From       string
	To         string
	Search     string
}

func GetQueryFilters(c *gin.Context) QueryFilters {
	pagination := GetPaginationParams(c)
	sort := c.DefaultQuery("sort", "desc")
	from := c.DefaultQuery("from", time.Now().Format("2006-01-01"))
	to := c.DefaultQuery("to", time.Now().AddDate(0, 0, 1).Format("2006-01-02"))
	search := c.DefaultQuery("search", "")

	return QueryFilters{
		Pagination: pagination,
		Sort:       sort,
		From:       from,
		To:         to,
		Search:     search,
	}
}
