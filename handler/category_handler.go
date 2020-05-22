package handler

import (
    "net/http"
    "github.com/labstack/echo"

    "../db"
	"../models"
)

// 存在するカテゴリーを返す
func  ShowCategorys() echo.HandlerFunc {
	return func(c echo.Context) error {
        categorys := []models.Category{}
        db.GetDB().Select("id, title").Find(&categorys)
        return c.JSON(http.StatusOK, categorys)
    }
}