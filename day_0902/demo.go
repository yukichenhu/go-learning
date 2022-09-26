package main

import (
	"github.com/cloudflare/tableflip"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

func subClient() {
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.String(200, "hello")
	})
	r.Run(":9091")
}

func main() {
	upg, _ := tableflip.New(tableflip.Options{})
	defer upg.Stop()
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT)
		for range sig {
			upg.Upgrade()
		}
	}()
	ln, _ := upg.Listen("tcp", "localhost:9090")
	defer ln.Close()
	r := gin.Default()
	r.GET("/index", func(context *gin.Context) {
		context.String(200, "index")
	})
	go subClient()
	go r.RunListener(ln)
	if err := upg.Ready(); err != nil {
		panic(err)
	}

	<-upg.Exit()
}
