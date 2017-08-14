package model

type Group struct {
    ID      uint        `gorm:"primary_key" json:"id" form:"id"`
    Name    string      `gorm:"primary_key" json:"name" form:"name"`
}

type Groups []Group
