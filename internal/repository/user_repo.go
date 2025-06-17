package repository

import (
	"CADET_PROJECT/internal/domain"

	"gorm.io/gorm"
)

type UserRepo struct{ db *gorm.DB }

func NewUserRepo(db *gorm.DB) *UserRepo { return &UserRepo{db} }

func (r *UserRepo) Create(u *domain.User) error { return r.db.Create(u).Error }
func (r *UserRepo) Find(id uint) (*domain.User, error) {
	var u domain.User
	return &u, r.db.First(&u, id).Error
}
func (r *UserRepo) List() ([]domain.User, error) { var v []domain.User; return v, r.db.Find(&v).Error }
func (r *UserRepo) Update(u *domain.User) error  { return r.db.Save(u).Error }
func (r *UserRepo) Delete(id uint) error         { return r.db.Delete(&domain.User{}, id).Error }
