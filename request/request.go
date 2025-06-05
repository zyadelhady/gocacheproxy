package request

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleReq(context *gin.Context) {
	url := context.Request.URL.Path
	fmt.Println(url)
	context.JSON(http.StatusOK, gin.H{
		"result": url,
	})
}
