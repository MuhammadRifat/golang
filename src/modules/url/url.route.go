package url

import (
	"url-shortner/src/util"

	"github.com/gin-gonic/gin"
)

func RegisterUrlRoutes(router *gin.Engine) {
	urlGroup := router.Group("/url")
	{
		urlGroup.POST("", util.JwtAuthMiddleware(), UrlController.CreateHandler)
		urlGroup.GET("", UrlController.FindAllHandler)
		urlGroup.GET("/by-id/:id", UrlController.FindOneHandler)
		urlGroup.GET("/:code", UrlController.RedirectHandler)
	}
}
