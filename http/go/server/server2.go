package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Server represents an HTTP server with a unique ID and port.
type Server struct {
	ID    string
	Port  string
	Mutex sync.Mutex
}

// NewServer creates a new Server instance.
func NewServer(id, port string) *Server {
	return &Server{
		ID:   id,
		Port: port,
	}
}

// Start initializes the HTTP server and begins listening for requests.
func (s *Server) Start() {
	http.HandleFunc("/hello", s.helloHandler)
	http.HandleFunc("/health", s.healthHandler)

	addr := fmt.Sprintf(":%s", s.Port)
	log.Printf("Server %s listening on %s\n", s.ID, addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server %s failed to start: %v\n", s.ID, err)
	}
}

// helloHandler responds to requests at the /hello endpoint.
func (s *Server) helloHandler(w http.ResponseWriter, r *http.Request) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	response := fmt.Sprintf("Hello from Server %s\n", s.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
	log.Printf("Server %s handled a request\n", s.ID)
}

// healthHandler responds to requests at the /health endpoint.
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK\n"))
	log.Printf("Server %s health check\n", s.ID)
}
