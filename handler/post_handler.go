package handler

import (
    "net/http"
    "github.com/labstack/echo"

    "../db"
	"../models"
)

// 全ての記事の(タイトル, 記事id)を返す
func ShowPosts() echo.HandlerFunc {
	return func(c echo.Context) error {
        posts := []models.Post{}
        db.GetDB().Select("id, title").Find(&posts)
        return c.JSON(http.StatusOK, posts)
    }
}

// 引数で指定されたidの記事の詳細情報を返す
func Detail() echo.HandlerFunc {
    return func(c echo.Context) error {
        detail := []models.Post{}
        db.GetDB().Where("id = ?", c.Param("post-id")).Find(&detail)
        return c.JSON(http.StatusOK, detail)
    }
}