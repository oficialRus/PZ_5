// internal/repository/product_repo.go
package repository

import (
	"CADET_PROJECT/internal/domain"

	"gorm.io/gorm"
)

type ProductRepo struct{ db *gorm.DB }

func NewProductRepo(db *gorm.DB) *ProductRepo { // ← именно ЭТОЙ сигнатуры ждёт main.go
	return &ProductRepo{db: db}
}

func (r *ProductRepo) Create(p *domain.Product) error { return r.db.Create(p).Error }
func (r *ProductRepo) Find(id uint) (*domain.Product, error) {
	var p domain.Product
	return &p, r.db.First(&p, id).Error
}
func (r *ProductRepo) List() ([]domain.Product, error) {
	var v []domain.Product
	return v, r.db.Find(&v).Error
}
func (r *ProductRepo) Update(p *domain.Product) error { return r.db.Save(p).Error }
func (r *ProductRepo) Delete(id uint) error           { return r.db.Delete(&domain.Product{}, id).Error }
