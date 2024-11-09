package storage

import (
	. "0x7266/preventives/core"
	"errors"
	"log"
)

type inMemoryDatabase struct {
	Technicians  *[]Technician
	Assets       *[]Asset
	Departments  *[]Department
	Maintenances *[]Maintenance
}

func NewInMemoryDatabase() *inMemoryDatabase {
	return &inMemoryDatabase{
		Technicians:  &[]Technician{},
		Assets:       &[]Asset{},
		Departments:  &[]Department{},
		Maintenances: &[]Maintenance{},
	}
}

func (i *inMemoryDatabase) Connect() {
	log.Println("In memory database created")
}

func (i *inMemoryDatabase) CreateTechnician(t *Technician) {
	*i.Technicians = append(*i.Technicians, *t)
}

func (i *inMemoryDatabase) GetTechnicians() []Technician {
	return *i.Technicians
}

func (i *inMemoryDatabase) GetTechnicianById(id string) (*Technician, error) {
	for index := range *i.Technicians {
		if (*i.Technicians)[index].Id == id {
			return &(*i.Technicians)[index], nil
		}
	}
	return nil, errors.New("Technician not found")
}

func (i *inMemoryDatabase) UpdateTechnician(t *Technician, p string) {
	panic("unimplemented")
}

func (i *inMemoryDatabase) DeleteTechnician(id string) error {
	for index := range *i.Technicians {
		if (*i.Technicians)[index].Id == id {
			*i.Technicians = append((*i.Technicians)[:index], (*i.Technicians)[index+1:]...)
			return nil
		}
	}
	return errors.New("Technician not found")
}
