package model

import (
    "github.com/jinzhu/gorm"
)

type Group struct {
    gorm.Model
    Name        string      `json:"name" form:"name"`
    Members     Users       `json:"members" form:"members"`
}

type Groups []Group
