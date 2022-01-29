package blog

import (
	"content/business/blog"

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

func (r Repository) FindContentById(id int) (*blog.Blog, error) {
	return &blog.Blog{}, nil
}
