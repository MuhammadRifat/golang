package url

import (
	"url-shortner/util"

	"github.com/gin-gonic/gin"
)

func RegisterUrlRoutes(router *gin.Engine) {
	urlGroup := router.Group("/url")
	{
		urlGroup.POST("", util.JwtAuthMiddleware(), CreateHandler)
		urlGroup.GET("", util.JwtAuthMiddleware(), FindAllHandler)
		urlGroup.GET("/by-id/:id", FindOneHandler)
		urlGroup.GET("/:code", RedirectHandler)
	}
}
