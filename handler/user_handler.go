package handler

import (
	"net/http"
	"time"
	
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"

	"../db"
	"../models"
)

type jwtCustomClaims struct {
    UserId  string    `json:"uid"`
    Name string `json:"name"`
    jwt.StandardClaims
}

var signingKey = []byte("secret-key")

// JWTによる認証でログイン状態を管理する
func Login(c echo.Context) error {
	userid := c.Param("user_id")
	password := c.Param("password")

	user := models.User{}
	db.GetDB().Where("user_id = ?", userid).Find(&user)

    if user.Password != password {
        return &echo.HTTPError{
            Code:    http.StatusUnauthorized,
            Message: "invalid name or password",
        }
	}

	claims :=  &jwtCustomClaims {
		user.UserId,
		user.Name,
        jwt.StandardClaims {
            ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
        },
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
    t, err := token.SignedString(signingKey)
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, map[string]string{
        "token": t,
    })
}
