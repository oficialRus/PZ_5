package domain

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"uniqueIndex;not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
}
