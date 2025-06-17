package service

import (
	"CADET_PROJECT/internal/domain"
	"CADET_PROJECT/internal/repository"

	"github.com/go-playground/validator/v10"
)

type (
	ProductService struct {
		repo *repository.ProductRepo
		v    *validator.Validate
	}
	ProductDTO struct {
		Code  string `json:"code"  validate:"required"`
		Name  string `json:"name"  validate:"required"`
		Price int64  `json:"price" validate:"required,gt=0"`
	}
)

func NewProductService(r *repository.ProductRepo) *ProductService {
	return &ProductService{repo: r, v: validator.New()}
}

func (s *ProductService) Create(dto ProductDTO) (*domain.Product, error) {
	if err := s.v.Struct(dto); err != nil {
		return nil, err
	}
	p := domain.Product{Code: dto.Code, Name: dto.Name, Price: dto.Price}
	if err := s.repo.Create(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (s *ProductService) Get(id uint) (*domain.Product, error) { return s.repo.Find(id) }
func (s *ProductService) List() ([]domain.Product, error)      { return s.repo.List() }
func (s *ProductService) Delete(id uint) error                 { return s.repo.Delete(id) }
