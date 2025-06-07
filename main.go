package main

import (
	"os"

	"gocacheproxy/ctx"
	"gocacheproxy/server"

	"github.com/gin-gonic/gin"
)

func main() {
	context := ctx.New()
	defer context.Cancel()
	url := os.Args[1]
	srv := server.New(url, context)
	r := gin.Default()
	srv.DataRoutes(r)
	r.NoRoute(srv.HandleReq)
	r.Run()
}
