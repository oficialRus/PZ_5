package models

import "fmt"

type Cadet struct {
	ID   int
	Name string
	Rank string
}

func NewCadet(id int, name, rank string) Cadet {
	return Cadet{ID: id, Name: name, Rank: rank}

}

func (c Cadet) String() string {
	return fmt.Sprintf("%d: %s (%s)", c.ID, c.Name, c.Rank)
}
