package model

import (
    "github.com/jinzhu/gorm"
)

type Group struct {
    gorm.Model
    Name        string      `json:"name" form:"name"`
}

type Groups []Group
