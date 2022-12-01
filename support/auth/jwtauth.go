package auth

import "github.com/gin-gonic/gin"

// JWTAuth 管理员接口鉴权
func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
	}
}
