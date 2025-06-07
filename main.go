package main

import (
	"os"

	"gocacheproxy/server"

	"github.com/gin-gonic/gin"
)

func main() {
	url := os.Args[1]
	srv := server.New(url)
	r := gin.Default()
	srv.DataRoutes(r)
	r.NoRoute(srv.HandleReq)
	r.Run()
}
