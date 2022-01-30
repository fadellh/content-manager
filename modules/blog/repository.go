package blog

import (
	"content/business/blog"
	"time"

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

func ToContentTable(b blog.Blog) *ContentTable {
	return &ContentTable{
		Title:       b.Title,
		Content:     b.Content,
		PublishedAt: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		CreatedAt:   time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}
}

type ContentTable struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement"`
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

func (r Repository) InsertContent(b blog.Blog) (int, error) {

	content := ToContentTable(b)

	err := r.DB.Create(&content).Error

	if err != nil {
		return 0, err
	}

	return content.ID, nil
}

func (r Repository) UpdateContent(b blog.Blog) error {

	var content ContentTable

	content.ID = b.ID

	tx := r.DB.Begin()

	err := tx.Model(&content).Updates(ContentTable{
		Title:     b.Title,
		Content:   b.Content,
		UpdatedAt: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Error; err != nil {
		return gorm.ErrRecordNotFound
	}

	return tx.Commit().Error
}
