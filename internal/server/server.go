package server

import (
	"fmt"
	"net/http"
)

// Server représente le serveur Go-net.
type Server struct {
	Port int
}

// NewServer crée une nouvelle instance de serveur.
func NewServer(port int) *Server {
	return &Server{
		Port: port,
	}
}

// Start démarre le serveur.
func (s *Server) Start() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go-net!")
	})

	fmt.Printf("Server is running on port %d...\n", s.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil)
}
