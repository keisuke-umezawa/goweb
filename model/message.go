package model

import (
    "github.com/jinzhu/gorm"
)

type Message struct {
    gorm.Model
    Sender      User
    GroupId     Group
    Text        string      `json:"text" form:"text"`
    Mode        string      `json:"mode" form:"mode"`
}

type Messages []Message



