package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"./db"
	"./handler"
	"./models"
)

func Test() {

	model := db.GetDB()

	model.Create(&models.Post{Title:"サンプル1", Text:"サンプルデータの2つ目です！", Image:"sample.com", UserId:1, Pv:0})
	model.Create(&models.Post{Title:"サンプル1", Text:"サンプルデータの3つ目です！", Image:"sample.com", UserId:2, Pv:0})
	model.Create(&models.Post{Title:"サンプル1", Text:"サンプルデータの4つ目です！", Image:"sample.com", UserId:3, Pv:0})

}

func main() {
	e := echo.New()

	// 利用するミドルウェアの宣言
    e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// DB接続
	db.Init()
	defer db.Close()

	Test()

	// ルーティング
	e.GET("/posts", handler.ShowPosts())
	e.GET("/detail/:post-id", handler.Detail())

	e.GET("/categorys", handler.ShowCategorys())
	
	e.Start(":8080")
}