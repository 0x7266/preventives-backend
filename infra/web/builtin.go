package web

import (
	"0x7266/preventives/core"
	"encoding/json"
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

type Technician struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (s *server) Run() error {
	s.storage.Connect()
	router := http.NewServeMux()
	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Ok"))
	})
	router.HandleFunc("GET /technicians", func(w http.ResponseWriter, r *http.Request) {
		log.Println(s.storage.GetTechnicians())
	})
	router.HandleFunc("GET /technicians/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		technician, err := s.storage.GetTechnicianById(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		data, err := json.Marshal(technician)
		if err != nil {
			w.WriteHeader(http.StatusTeapot)
			w.Write([]byte("Something went wrong while parsing the data"))
		}
		w.Write(data)

	})
	router.HandleFunc("POST /technicians", func(w http.ResponseWriter, r *http.Request) {
		log.Println("irstnirsnie")
		var technician Technician
		json.NewDecoder(r.Body).Decode(&technician)
		s.storage.CreateTechnician(&core.Technician{
			Id:   technician.Id,
			Name: technician.Name,
		})
	})

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
