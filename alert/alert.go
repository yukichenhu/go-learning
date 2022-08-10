package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learning/entity"
	"net/http"
)

var r *gin.Engine

func main() {
	r = gin.Default()
	r.POST("/webhook", func(context *gin.Context) {
		var notify entity.Notification
		err := context.BindJSON(&notify)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		fmt.Println("报警接受成功", notify)
		context.JSON(http.StatusOK, "receive success")
	})
	r.POST("/addMapping", func(context *gin.Context) {
		method := context.PostForm("method")
		url := context.PostForm("url")
		resp := context.PostForm("resp")
		fmt.Println(fmt.Sprintf("method:%s,url:%s,resp:%s", method, url, resp))
		addMapping(method, url, resp)
		context.String(200, "add success")
	})
	r.Run(":8888")
}

func addMapping(method string, url string, resp string) {
	switch method {
	case "POST":
		r.POST(url, func(context *gin.Context) {
			context.JSON(200, gin.H{"msg": resp})
		})
	case "GET":
		r.GET(url, func(context *gin.Context) {
			context.JSON(200, gin.H{"msg": resp})
		})
	default:
		r.Any(url, func(context *gin.Context) {
			context.JSON(200, gin.H{"msg": resp})
		})
	}
}
