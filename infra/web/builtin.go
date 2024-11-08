package web

import (
	"0x7266/preventives/core"
	"log"
	"net/http"
)

type server struct {
	address string
	storage core.Storage
}

func NewBuiltInServer(address string, storage core.Storage) *server {
	return &server{
		address,
		storage,
	}
}

func (s *server) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Ok"))
	})
	router.HandleFunc("GET /technicians", func(w http.ResponseWriter, r *http.Request) {
		s.storage.GetTechnician()
	})
	// v1 := http.NewServeMux()
	// v1.Handle("/api/v1", router)

	server := http.Server{
		Addr:    s.address,
		Handler: router,
	}

	log.Printf(
		"Server running on %s",
		s.address,
	)
	return server.ListenAndServe()
}
