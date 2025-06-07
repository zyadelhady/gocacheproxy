package request

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleReq(context *gin.Context, url string) {
	path := context.Request.URL.Path
	fullUrl := url + path
	context.JSON(http.StatusOK, gin.H{
		"result": fullUrl,
	})
}
