package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"short-url/support/rsa"
)

// Login 后台管理员登录
func Login(c *gin.Context) {
	param := make(map[string]interface{})
	err := c.BindJSON(&param)
	if err != nil {
		return
	}
	err, pwd := rsa.PrivateKeyDecrypt(param["password"].(string))
	if err != nil {
		fmt.Println("密文：" + param["password"].(string))
		c.JSON(http.StatusOK, gin.H{"password": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"password": string(pwd)})
	}
}
