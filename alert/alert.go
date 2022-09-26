package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.String(200, "hello "+strconv.Itoa(os.Getpid()))
	})
	go openChild()
	_ = endless.ListenAndServe(":8888", r)
}

func openChild() {
	r := gin.Default()
	r.GET("/child", func(context *gin.Context) {
		context.String(200, "child "+strconv.Itoa(os.Getpid()))
	})
	_ = endless.ListenAndServe(":8889", r)
}
