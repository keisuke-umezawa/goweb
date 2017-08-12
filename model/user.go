package model

type User struct {
    Id      uint        `gorm:"primary_key;Auto_INCREMENT" json:"id" form:"id"`
    Name    string      `json:"name" form:"name"`
}

type Users []User

