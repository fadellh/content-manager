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

func (c ContentTable) ToBlogDomain() *blog.Blog {

	return &blog.Blog{
		ID:          c.ID,
		Title:       c.Title,
		Content:     c.Content,
		PublishedAt: c.PublishedAt,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}

}

type ContentTable struct {
	ID          int    `gorm:"column:id;primaryKey"`
	Content     string `gorm:"column:content"`
	Title       string `gorm:"column:title"`
	PublishedAt string `gorm:"column:published_at"`
	CreatedAt   string `gorm:"column:created_at"`
	UpdatedAt   string `gorm:"column:updated_at"`
}

func (r Repository) FindContentById(id int) (*blog.Blog, error) {

	var content ContentTable

	err := r.DB.Where("id = ?", id).Find(&content).Error

	if err != nil {
		return nil, err
	}

	return content.ToBlogDomain(), nil
}

func (r Repository) FindAllContent() ([]blog.Blog, error) {

	var contents []ContentTable
	var blogs []blog.Blog

	err := r.DB.Find(&contents).Error

	if err != nil {
		return nil, err
	}

	for _, item := range contents {
		blogs = append(blogs, *item.ToBlogDomain())
	}

	return blogs, nil
}
