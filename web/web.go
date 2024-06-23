package web

import (
	"log"
	"net/http"
	"os"

	"github.com/RaghavSood/spoonfeed/feedgen"
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
	log.Println("Starting server on port", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Println("Error starting server:", err)
	}
}

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (s *Server) RSS(w http.ResponseWriter, r *http.Request) {
	feedGen := feedgen.NewFeedGenerator(feedgen.RSS2)
	feed, err := feedGen.Generate()
	if err != nil {
		w.Write([]byte("Error generating feed"))
		return
	}

	w.Header().Set("Content-Type", "application/rss+xml")
	w.Write(feed.EncodedContent)
}

func (s *Server) RSSBucket(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("RSS Feed for bucket"))
}
