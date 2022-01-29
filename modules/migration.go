package modules

import (
	"content/modules/blog"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(blog.ContentTable{})

}
