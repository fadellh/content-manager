package blog

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewGormDB(db *gorm.DB) *Repository {
	return &Repository{
		db,
	}
}
