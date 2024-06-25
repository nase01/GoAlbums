package handlers

import (
	"net/http"

	"GoAlbums/internal/service"
	"GoAlbums/utils/helpers"

	"github.com/gin-gonic/gin"
)

func GetUserLogs(c *gin.Context) {

	pagination := helpers.GetPaginationParams(c)

	userLogs, err := service.GetUserLogs(pagination.CurrentPage, pagination.PerPage)
	if err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}
	c.JSON(http.StatusOK, userLogs)
}
