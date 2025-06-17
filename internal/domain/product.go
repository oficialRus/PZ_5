package domain

type Product struct {
	ID    uint   `gorm:"primaryKey"`
	Code  string `gorm:"uniqueIndex;not null"`
	Name  string `gorm:"not null"`
	Price int64  `gorm:"not null"` // храним в центах/копейках
}

func (Product) TableName() string { return "products" } // необязательно, но так явно
