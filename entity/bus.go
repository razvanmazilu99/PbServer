package entity

type Bus struct {
	Name  string
	Model int
	Type  string
}

func (bus Bus) CanDrive() bool {
	return false
}
