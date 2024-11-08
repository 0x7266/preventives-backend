package core

type Technician struct {
	Name string
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
