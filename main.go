package main

import (
	"log"
	"net/http"
)

type Server struct {
	address  string
	database Database
}

func NewServer(address string, database Database) *Server {
	return &Server{
		address,
		database,
	}
}

func (s *Server) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Ok"))
	})
	router.HandleFunc("GET /technicians", func(w http.ResponseWriter, r *http.Request) {
		s.database.getTechnician()
	})
	// v1 := http.NewServeMux()
	// v1.Handle("/api/v1", router)

	server := http.Server{
		Addr:    s.address,
		Handler: router,
	}

	log.Printf("Server running on %s", s.address)
	return server.ListenAndServe()
}

type Database interface {
	connect()
	createTechnician(t *Technician)
	getTechnician()
	updateTechnician(t *Technician, p string)
	deleteTechnician(t *Technician)
}

type Technician struct {
	name string
}

type Asset struct {
	name       string
	department Department
}

type Department struct {
	name   string
	assets []Asset
}

type Maintenance struct {
	date       string
	state      string
	department *Department
	asset      *Asset
}

type InMemoryDatabase struct {
	technicians  *[]Technician
	assets       *[]Asset
	departments  *[]Department
	maintenances *[]Maintenance
}

func (i *InMemoryDatabase) connect() {
	panic("unimplemented")
}

func (i *InMemoryDatabase) createTechnician(t *Technician) {
	panic("unimplemented")
}

func (i *InMemoryDatabase) deleteTechnician(t *Technician) {
	panic("unimplemented")
}

func (i *InMemoryDatabase) getTechnician() {
	log.Println(i.technicians)
}

func (i *InMemoryDatabase) updateTechnician(t *Technician, p string) {
	panic("unimplemented")
}

func main() {
	server := NewServer(":3333", &InMemoryDatabase{
		technicians: &[]Technician{
			{name: "aaaaaaaaaaaa"},
			{name: "bbbbbbbbbbbb"},
			{name: "ccccccccccc"},
		},
		assets:       &[]Asset{},
		departments:  &[]Department{},
		maintenances: &[]Maintenance{},
	})
	log.Fatalln(server.Run())
}
