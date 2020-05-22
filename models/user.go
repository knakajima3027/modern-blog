package models

type User struct {
    UserId string `gorm:"primary_key" column:"user_id"`
    Name string `gorm:"column:name"`
	Password string `gorm:"column:password"`
}