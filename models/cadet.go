package models

import "fmt"

type Cadet struct {
	Name string
	Role string
}

func NewCadet(name, role string) Cadet {
	return Cadet{Name: name, Role: role}
}
func (c Cadet) String() string {
	return fmt.Sprintf("%s (%s)", c.Name, c.Role)
}
