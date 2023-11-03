package config

import (
	"GO_API_BOILERPLATE/users"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	PORT string = ":8080"
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
		Addr:           PORT,
		Handler:        s.server,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	return err
}
