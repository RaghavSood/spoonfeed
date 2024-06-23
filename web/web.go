package web

import (
	"net/http"
	"os"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Serve() {
	router := http.NewServeMux()
	router.HandleFunc("GET /", s.Index)
	router.HandleFunc("GET /rss2", s.RSS)
	router.HandleFunc("GET /rss2/{bucket}", s.RSSBucket)

	port := os.Getenv("SPOONFEED_PORT")
	http.ListenAndServe(":"+port, router)
}

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (s *Server) RSS(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("RSS Feed"))
}

func (s *Server) RSSBucket(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("RSS Feed for bucket"))
}
