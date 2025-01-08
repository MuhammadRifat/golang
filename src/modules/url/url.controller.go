package url

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UrlControllerStruct struct{}

var UrlController = UrlControllerStruct{}

func (c *UrlControllerStruct) CreateHandler(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	var urlDto URLDto
	if err := ctx.ShouldBindJSON(&urlDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userIdInt, ok := userId.(int)
	if !ok {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userId is not a string"})
		return
	}

	urlDto.UserId = userIdInt
	user, err := UrlService.Create(urlDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UrlControllerStruct) FindOneHandler(ctx *gin.Context) {
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

func (c *UrlControllerStruct) RedirectHandler(ctx *gin.Context) {
	code := ctx.Param("code")

	url, err := UrlService.FindOneByCode(code)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Redirect(http.StatusFound, url.OriginalUrl)
}

func (c *UrlControllerStruct) FindAllHandler(ctx *gin.Context) {
	urls, err := UrlService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, urls)
}
