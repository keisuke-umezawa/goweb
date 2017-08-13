package model

import (
    "github.com/jinzhu/gorm"
)

type Message struct {
    gorm.Model
    Sender      User        `gorm:"ForeignKey:SenderID"`
    SenderID    uint
    Group       Group
    GroupID     uint
    Text        string      `json:"text" form:"text"`
    Mode        string      `json:"mode" form:"mode"`
}

type Messages []Message



