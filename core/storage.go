package core

type Storage interface {
	Connect()
	CreateTechnician(t *Technician)
	GetTechnician()
	UpdateTechnician(t *Technician, p string)
	DeleteTechnician(t *Technician)
}
