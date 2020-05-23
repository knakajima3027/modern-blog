package handler

import (
    "net/http"
    
    "github.com/labstack/echo"
    "github.com/gorilla/sessions"
    "github.com/labstack/echo-contrib/session"
	"../db"
    "../models"
    "fmt"
)

func Login() echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.FormValue("userid")
        password := c.FormValue("password")

        user := models.User{}
        db.GetDB().Where("user_id = ?", userid).Find(&user)

        if user.Password != password {
            return &echo.HTTPError{
                Code:    http.StatusUnauthorized,
                Message: "invalid name or password",
            }
        }
        
        sess, _ := session.Get("session", c)

        fmt.Println(sess)

        sess.Options = &sessions.Options{
            MaxAge:   86400 * 3, // cookie有効期限 (3日)
            HttpOnly: true,
        }

        sess.Values["auth"] = true

        if err:=sess.Save(c.Request(), c.Response());err!=nil{
            return c.NoContent(http.StatusInternalServerError)
        }

        return c.NoContent(http.StatusOK)
    }
}

func Logout() echo.HandlerFunc {
    return func(c echo.Context)error { 
        sess, _ := session.Get("session", c)
        sess.Values["auth"]=false
        if err:=sess.Save(c.Request(), c.Response());err!=nil{
            return c.NoContent(http.StatusInternalServerError)
        }
        return c.NoContent(http.StatusOK)
    }
}

// ログイン状態なら「hey」と返す
func Secret() echo.HandlerFunc {
    return func(c echo.Context)error {
        
        sess, err := session.Get("session", c)
        if err!=nil {
            fmt.Println(err)
            return c.String(http.StatusInternalServerError, "Error")
        }
        
        if b, _:=sess.Values["auth"];b!=true{
            return c.String(http.StatusUnauthorized, "401")
        }else {
            return c.String(http.StatusOK, "hey!")
        }
    }
}