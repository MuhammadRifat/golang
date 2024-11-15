package url

import "github.com/gin-gonic/gin"

func RegisterUrlRoutes(router *gin.Engine) {
	urlGroup := router.Group("/url")
	{
		urlGroup.POST("", CreateHandler)
		urlGroup.GET("", FindAllHandler)
		urlGroup.GET("/:id", FindOneHandler)
	}
}
