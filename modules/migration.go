package modules

import (
	"content/modules/blog"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	err := db.AutoMigrate(blog.ContentTable{})

	if err == nil && db.Migrator().HasTable(&blog.ContentTable{}) {
		err = db.FirstOrCreate(&blog.ContentTable{
			ID:          1,
			Title:       "Hello world",
			Content:     "Hello world dang dang",
			PublishedAt: "2021-10-13T05:07:57.208Z",
			CreatedAt:   "2021-10-13T05:07:57.208Z",
			UpdatedAt:   "2021-10-13T05:07:57.208Z",
		}).Error
		err = db.FirstOrCreate(&blog.ContentTable{
			ID:          2,
			Title:       "Number 2",
			Content:     "Number 2 dang dang",
			PublishedAt: "2021-10-13T05:07:57.208Z",
			CreatedAt:   "2021-10-13T05:07:57.208Z",
			UpdatedAt:   "2021-10-13T05:07:57.208Z",
		}).Error
	}

}
