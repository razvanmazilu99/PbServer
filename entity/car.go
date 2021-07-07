package entity

type Car struct {
	Name  string `json:"name"`
	Model int    `json:"model"`
	Type  string `json:"type"`
}

func (car Car) CanDrive() bool {
	return true
}
