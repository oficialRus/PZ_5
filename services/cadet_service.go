package services

import (
	"cadet_project/config"
	"cadet_project/models"
)

type CadetService struct {
	cadet models.Cadet
}

// Конструктор
func NewCadetService() *CadetService {
	name := config.Get("CADET_NAME", "Ruslan")
	role := config.Get("CADET_ROLE", "Cadet")
	return &CadetService{
		cadet: models.NewCadet(name, role),
	}

}

func (s *CadetService) GetCadetInfo() string {
	return s.cadet.String()
}
