package url

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateHandler(ctx *gin.Context) {
	var urlDto URLDto
	if err := ctx.ShouldBindJSON(&urlDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := UrlService.Create(urlDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func FindOneHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	url, err := UrlService.FindOneById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, url)
}

func FindAllHandler(ctx *gin.Context) {
	urls, err := UrlService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, urls)
}
