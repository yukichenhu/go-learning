package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Group("v1").GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "pong",
		})
	})
	r.Run(":8888")
}

func method1() {

}
