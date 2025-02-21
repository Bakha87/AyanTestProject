package server

import (
	"AyanTestProject/internal/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	handler *handler.Handler
	srv     *gin.Engine
}

func New(h *handler.Handler) *Server {
	return &Server{
		handler: h,
	}
}

func (s *Server) Run() {
	s.srv = gin.Default()
	s.srv.Use(gin.Recovery())

	port := ":8080"

	Endpoint(s, "v1")

	srv := &http.Server{
		Addr:           port,
		Handler:        s.srv,
		ReadTimeout:    300 * time.Second,
		WriteTimeout:   300 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
