package modules

import (
	"content/modules/blog"
	"time"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	err := db.AutoMigrate(blog.ContentTable{})

	if err == nil && db.Migrator().HasTable(&blog.ContentTable{}) {
		err = db.FirstOrCreate(&blog.ContentTable{
			ID:          1,
			Title:       "Hello world",
			Content:     "Hello world dang dang",
			PublishedAt: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
			CreatedAt:   time.Now().UTC().Format("2006-01-02T15:04:05Z"),
			UpdatedAt:   time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		}).Error
		err = db.FirstOrCreate(&blog.ContentTable{
			ID:          2,
			Title:       "Number 2",
			Content:     "Number 2 dang dang",
			PublishedAt: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
			CreatedAt:   time.Now().UTC().Format("2006-01-02T15:04:05Z"),
			UpdatedAt:   time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		}).Error
	}

}
