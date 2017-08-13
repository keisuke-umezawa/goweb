package model

import (
    "github.com/jinzhu/gorm"
)

type Message struct {
    gorm.Model
    SenderId    uint        `json:"sender_id" form:"sender_id"`
    GroupId     uint        `json:"group_id" form:"group_id"`
    Text        string      `json:"text" form:"text"`
    Mode        string      `json:"mode" form:"mode"`
}

type Messages []Message



