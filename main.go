package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	"./db"
	"./handler"
	"./models"
)

func Test() {

	model := db.GetDB()

	model.Create(&models.Post{Title:"サンプル1", Text:"サンプルデータの2つ目です！", Image:"sample.com", UserId:1})
	model.Create(&models.Post{Title:"サンプル1", Text:"サンプルデータの3つ目です！", Image:"sample.com", UserId:2})
	model.Create(&models.Post{Title:"サンプル1", Text:"サンプルデータの4つ目です！", Image:"sample.com", UserId:3})

}

func CreateAdmin() {
	model := db.GetDB()
	model.Create(&models.User{UserId:"sample", Name:"kouhei", Password:"password"})
}

func main() {
	e := echo.New()

	// 利用するミドルウェアの宣言
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	// DB接続
	db.Init()
	defer db.Close()

	//Test()
	//CreateAdmin()

	// ルーティング

	//管理ユーザー関連	
	e.POST("/login", handler.Login())
	e.GET("/secret", handler.Secret())

	// 記事関連
	e.GET("/posts", handler.ShowPosts())
	e.GET("/detail/:post-id", handler.Detail())

	// カテゴリー関連
	e.GET("/categorys", handler.ShowCategorys())
	
	e.Start(":8080")
}