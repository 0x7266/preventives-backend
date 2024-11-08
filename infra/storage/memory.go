package storage

import (
	. "0x7266/preventives/core"
	"log"
)

type InMemoryDatabase struct {
	Technicians  *[]Technician
	Assets       *[]Asset
	Departments  *[]Department
	Maintenances *[]Maintenance
}

func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{
		Technicians:  &[]Technician{},
		Assets:       &[]Asset{},
		Departments:  &[]Department{},
		Maintenances: &[]Maintenance{},
	}
}

func (i *InMemoryDatabase) Connect() {
	panic("unimplemented")
}

func (i *InMemoryDatabase) CreateTechnician(t *Technician) {
	panic("unimplemented")
}

func (i *InMemoryDatabase) DeleteTechnician(t *Technician) {
	panic("unimplemented")
}

func (i *InMemoryDatabase) GetTechnician() {
	log.Println(i.Technicians)
}

func (i *InMemoryDatabase) UpdateTechnician(t *Technician, p string) {
	panic("unimplemented")
}
