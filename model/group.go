package model

import _ "github.com/jinzhu/gorm"

type Group struct {
    ID      uint        `gorm:"primary_key" json:"id" form:"id"`
    Name    string      `gorm:"unique" json:"name" form:"name"`
}

type Groups []Group
