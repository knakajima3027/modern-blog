package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"../models"
)

var (
	db *gorm.DB
	err error
)

func Init() {
	// DBにコネクト
	db, err = gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		panic(err)
	}

	// ログを表示する
	db.LogMode(true)

	// テーブルが存在しない場合に自動でテーブルを作成
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Category{})
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}