package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Title string `gorm:"column:title"`
	Text string `gorm:"column:text"`
	Image string `gorm:"column:image"`
	UserId uint `gorm:"column:user_id"`
}