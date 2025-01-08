package url

import (
	"net/http"
	"strconv"
	"url-shortner/src/util"

	"github.com/gin-gonic/gin"
)

type UrlControllerStruct struct{}

var UrlController = UrlControllerStruct{}

func (c *UrlControllerStruct) CreateHandler(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.Error(util.UnauthorizedErr())
		return
	}

	var urlDto URLDto
	if err := ctx.ShouldBindJSON(&urlDto); err != nil {
		ctx.Error(util.ValidationErr(err))
		return
	}
	userIdInt, ok := userId.(int)
	if !ok {
		ctx.Error(util.BadRequestErr("userId must be a integer"))
		return
	}

	urlDto.UserID = userIdInt
	user, err := UrlService.CreateUrl(urlDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.AbortWithStatusJSON(util.ResponseCreated(user))
}

func (c *UrlControllerStruct) FindOneHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	url, err := UrlService.FindUrlById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, url)
}

func (c *UrlControllerStruct) RedirectHandler(ctx *gin.Context) {
	code := ctx.Param("code")

	url, err := UrlService.FindUrlByCode(code)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Redirect(http.StatusFound, url.OriginalUrl)
}

func (c *UrlControllerStruct) FindAllHandler(ctx *gin.Context) {
	urls, err := UrlService.FindAllUrls()
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, urls)
}
