package services

import (
	"errors"

	"cadet_project/models"
	"cadet_project/store"
)

// EquipmentService реализует CRUD поверх общего in-memory хранилища.
type EquipmentService struct {
	mem *store.Memory[models.Equipment]
}

func NewEquipmentService() *EquipmentService {
	return &EquipmentService{
		mem: store.NewMemory[models.Equipment](),
	}
}

func (s *EquipmentService) List() []models.Equipment {
	return s.mem.All()
}

func (s *EquipmentService) Add(e models.Equipment) models.Equipment {
	// назначаем ID перед сохранением
	setID := func(ptr *models.Equipment, id int) { ptr.ID = id }
	return s.mem.WithID(setID)(e)
}

func (s *EquipmentService) Delete(id int) error {
	if ok := s.mem.Delete(func(e models.Equipment) bool { return e.ID == id }); !ok {
		return errors.New("equipment not found")
	}
	return nil
}

func (s *EquipmentService) Update(id int, in models.Equipment) error {
	ok := s.mem.Update(func(e *models.Equipment) bool {
		if e.ID == id {
			e.Name = in.Name
			e.Status = in.Status
			return true
		}
		return false
	})
	if !ok {
		return errors.New("equipment not found")
	}
	return nil
}
