package core

type Storage interface {
	Connect()
	CreateTechnician(t *Technician)
	GetTechnicians() []Technician
	GetTechnicianById(id string) (*Technician, error)
	UpdateTechnician(t *Technician, p string)
	DeleteTechnician(id string) error
}
