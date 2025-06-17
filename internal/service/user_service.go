package service

import (
	"CADET_PROJECT/internal/domain"
	"CADET_PROJECT/internal/repository"

	"github.com/go-playground/validator/v10"
)

type (
	UserService struct {
		repo *repository.UserRepo
		v    *validator.Validate
	}
	SignUpDTO struct {
		Email     string `json:"email" validate:"required,email"`
		FirstName string `json:"first_name" validate:"required"`
		LastName  string `json:"last_name"  validate:"required"`
	}
)

func NewUserService(r *repository.UserRepo) *UserService {
	return &UserService{repo: r, v: validator.New()}
}

func (s *UserService) Create(dto SignUpDTO) (*domain.User, error) {
	if err := s.v.Struct(dto); err != nil {
		return nil, err
	}
	u := domain.User{
		Email:     dto.Email,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	}
	if err := s.repo.Create(&u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *UserService) Get(id uint) (*domain.User, error) { return s.repo.Find(id) }
func (s *UserService) List() ([]domain.User, error)      { return s.repo.List() }
func (s *UserService) Delete(id uint) error              { return s.repo.Delete(id) }
