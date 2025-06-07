package server

import (
	"gocacheproxy/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) DataRoutes(engine *gin.Engine) {
	group := engine.Group("data")
	group.GET("/", getData)
}

func getData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "here is your data",
		"data":    []int{},
	})
}
