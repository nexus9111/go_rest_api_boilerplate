package config

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nexus9111/go-rest-api-boilerplate/users"
	"net/http"
	"os"
	"time"
)

type Server struct {
	server *gin.Engine
}

func (s *Server) Init() {
	r := gin.Default()
	r.Use(cors.Default())

	users.InitRouter(r)

	s.server = r
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		Handler:        s.server,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	return err
}
