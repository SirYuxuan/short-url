package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"net/http"
	"short-url/support/redispool"
)

// Visit 用户访问的主服务
func Visit(c *gin.Context) {
	key := c.Param("key")

	conn := redispool.Pool.Get()

	defer func(conn redis.Conn) {
		_ = conn.Close()
	}(conn)

	exists, _ := conn.Do("EXISTS", key)

	flag, _ := redis.Int(exists, nil)

	if flag == 1 {
		// 这里访问成功了，可以记录链接访问记录
		url, _ := conn.Do("GET", key)
		urlPath, _ := redis.String(url, nil)
		c.Redirect(http.StatusFound, urlPath)
	} else {
		c.Redirect(http.StatusFound, "https://www.yuxuan66.com")
	}

}
