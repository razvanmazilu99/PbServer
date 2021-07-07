package entity

type Bike struct {
	Name  string
	Model int
	Type  string
}

func (bike Bike) CanDrive() bool {
	return true
}
