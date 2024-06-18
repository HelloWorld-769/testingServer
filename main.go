package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()

	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Server Listening",
		})
	})

	engine.Run(":8080")

}
