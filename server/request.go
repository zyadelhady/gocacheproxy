package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CachedResponse struct {
	Headers map[string][]string `json:"headers"`
	Body    []byte              `json:"body"`
	Status  int                 `json:"status"`
}

func (s *Server) HandleReq(context *gin.Context) {
	path := context.Request.URL.Path
	w := context.Writer
	fullUrl := s.url + path

	redis := s.ctx.Redis()

	cachedRes, err := redis.Get(context, path).Result()
	if err == nil {
		log.Println("CACHE HIT")
		var cached CachedResponse
		if err := json.Unmarshal([]byte(cachedRes), &cached); err != nil {
			log.Printf("Failed to unmarshal cached data: %v", err)
			context.Status(http.StatusInternalServerError)
			return
		}
		for key, values := range cached.Headers {
			for _, value := range values {
				context.Writer.Header().Add(key, value)
			}
		}
		w.WriteHeader(cached.Status)
		w.Write(cached.Body)
		return
	}
	log.Println("CACHE MISS")
	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		context.Status(http.StatusInternalServerError)
		return
	}
	cached := CachedResponse{
		Headers: resp.Header,
		Body:    body,
		Status:  resp.StatusCode,
	}

	jsonBytes, _ := json.Marshal(cached)

	if err := redis.Set(context, path, jsonBytes, 5*time.Minute).Err(); err != nil {
		log.Printf("Error setting cache: %v", err)
	}

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}
