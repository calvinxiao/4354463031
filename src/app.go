package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var flag1, flag2 string

func init() {
	flag1 = os.Getenv("flag1")
	flag2 = os.Getenv("flag2")
}

func main() {
	r := gin.New()

	r.Use(func(ctx *gin.Context) {
		ctx.Next()
		method := ctx.Request.Method
		path := ctx.Request.URL.Path
		log.Println(fmt.Sprintf("%s %s", method, path))
	})

	r.GET("/users/:userID", func(c *gin.Context) {
		if c.Param("userID") != "1" {
			c.AbortWithStatus(404)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"name": "Mojito",
			},
		})
	})

	r.GET("/flag2", func(c *gin.Context) {
		headerApiKey := c.GetHeader("api-key")
		if len(headerApiKey) == 0 || flag1 != headerApiKey {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Data(http.StatusOK, "text/plain", []byte(flag2))
	})
	r.Run(":80")
}
