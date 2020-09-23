package router

import "github.com/gin-gonic/gin"

func InitRouters() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.GET("/user/info", func(context *gin.Context) {
		context.String(200, "this user info")
	})
	return ginRouter
}
