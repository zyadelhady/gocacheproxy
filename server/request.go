package server

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) HandleReq(context *gin.Context) {
	path := context.Request.URL.Path
	fullUrl := s.url + path
	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	context.Writer.WriteHeader(resp.StatusCode)

	for key, values := range resp.Header {
		for _, value := range values {
			context.Writer.Header().Add(key, value)
		}
	}
	io.Copy(context.Writer, resp.Body)
}
