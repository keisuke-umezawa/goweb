package model

type User struct {
    ID      uint        `json:"id" form:"id"`
    Name    string      `json:"name" form:"name"`
}

type Users []User


