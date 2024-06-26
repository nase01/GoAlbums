package handlers

import (
	"net/http"
	"time"

	"GoAlbums/internal/service"
	"GoAlbums/utils/helpers"

	"github.com/gin-gonic/gin"
)

func GetUserLogs(c *gin.Context) {

	pagination := helpers.GetPaginationParams(c)
	sort := c.DefaultQuery("sort", "desc")
	from := c.DefaultQuery("from", time.Now().Format("2006-01-01"))
	to := c.DefaultQuery("to", time.Now().AddDate(0, 0, 1).Format("2006-01-02"))

	userLogs, err := service.GetUserLogs(pagination.CurrentPage, pagination.PerPage, sort, from, to)
	if err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}
	c.JSON(http.StatusOK, userLogs)
}
