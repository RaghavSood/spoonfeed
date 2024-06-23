package web

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Serve() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", s.Index)

	port := os.Getenv("SPOONFEED_PORT")
	r.Run(":" + port)
}

func (s *Server) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}
