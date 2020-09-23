package router

import (
	"github.com/gin-gonic/gin"
	"go-micro-learn/members/handler/web"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.GET("/user/test", func(context *gin.Context) {
		context.String(200, "this user test")
	})
	router.POST("/user/login", web.Login)
	router.GET("/user/info", web.UserInfo)
	return router
}
