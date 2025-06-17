package services

import (
	"errors"

	"cadet_project/models"
	"cadet_project/store"
)

type CadetService struct{ mem *store.Memory[models.Cadet] }

func NewCadetService() *CadetService {
	return &CadetService{mem: store.NewMemory[models.Cadet]()}
}

func (s *CadetService) List() []models.Cadet { return s.mem.All() }

func (s *CadetService) Add(c models.Cadet) models.Cadet {
	setID := func(ptr *models.Cadet, id int) { ptr.ID = id }
	return s.mem.WithID(setID)(c)
}

func (s *CadetService) Delete(id int) error {
	if ok := s.mem.Delete(func(c models.Cadet) bool { return c.ID == id }); !ok {
		return errors.New("cadet not found")
	}
	return nil
}

func (s *CadetService) Update(id int, in models.Cadet) error {
	ok := s.mem.Update(func(c *models.Cadet) bool {
		if c.ID == id {
			c.Name = in.Name
			c.Rank = in.Rank
			return true
		}
		return false
	})
	if !ok {
		return errors.New("cadet not found")
	}
	return nil
}
