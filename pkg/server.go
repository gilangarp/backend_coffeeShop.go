package pkg

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Server(router *gin.Engine) *http.Server {
	addr := "0.0.0.0:8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}
	server := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 10,
		Handler:      router,
	}
	return server
}
