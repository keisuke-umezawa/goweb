package model

type User struct {
    ID      uint        `gorm:"primary_key" json:"id" form:"id"`
    Name    string      `gorm:"primary_key" json:"name" form:"name"`
}

type Users []User


