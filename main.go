package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// new 一个 Gin 实例
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {

		// 以 JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})
	r.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义， 请确认URL和请求是否正确。",
			})
		}
	})

	// 运行服务
	r.Run(":8888")
}
