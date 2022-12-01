package router

import (
	"github.com/gin-gonic/gin"
	"short-url/api/admin"
	"short-url/support/cors"
)
import "short-url/api"

func InitRouter() *gin.Engine {
	Router := gin.Default()
	// 跨域处理
	Router.Use(cors.Cors())
	// 访问此接口进行跳转
	Router.GET("/:key", api.Visit)
	// 管理后台接口
	Auth := Router.Group("auth")
	Auth.POST("/login", admin.Login)

	return Router
}
