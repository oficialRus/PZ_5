package store

import (
	"cadet_project/models"
	"errors"
	"sync"
)

type MemoryStore struct {
	mtx    sync.RWMutex
	cadets []models.Cadet
	nextID int
}

// MemoryStore хранит кадетов в срезе и защищён мьютексом.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		cadets: make([]models.Cadet, 0),
		nextID: 1,
	}
}

// /Метод отображения списка сущностей
func (s *MemoryStore) GetAll() []models.Cadet {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	out := make([]models.Cadet, len(s.cadets))
	copy(out, s.cadets)
	return out
}

// Добавление сущностей
func (s *MemoryStore) Add(name, rank string) models.Cadet {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	c := models.NewCadet(s.nextID, name, rank)
	s.nextID++
	s.cadets = append(s.cadets, c)
	return c
}

// /Удаляет
func (s *MemoryStore) Delete(id int) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	for i, c := range s.cadets {
		if c.ID == id {
			s.cadets = append(s.cadets[:i], s.cadets[i+1:]...)
			return nil
		}
	}

	return errors.New("cadet not found")
}

// Обновляет
func (s *MemoryStore) Update(id int, name, rank string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	for i, c := range s.cadets {
		if c.ID == id {
			s.cadets[i].Name = name
			s.cadets[i].Rank = rank
			return nil
		}
	}
	return errors.New("cadet not found ")
}
