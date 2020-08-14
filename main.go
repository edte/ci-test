// @program: ci-test
// @author: edte
// @create: 2020-08-13 20:58
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello world",
		})
	})

	r.Run()
}
