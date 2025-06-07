package main

import (
	"fmt"
	"os"

	"gocacheproxy/request"
	"gocacheproxy/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := os.Args[1]
	fmt.Println("server", server)
	r := gin.Default()
	routes.DataRoutes(r)
	r.NoRoute(func(ctx *gin.Context) {
		request.HandleReq(ctx,server)
	})
	r.Run()
}
