package model

import _ "github.com/jinzhu/gorm"

type Message struct {
    ID          uint        `json:"id" form:"id"`
    UserID      uint        `json:"user_id" form:"user_id"`
    GroupID     uint        `json:"group_id" form:"group_id"`
    Text        string      `json:"text" form:"text"`
    Mode        string      `json:"mode" form:"mode"`
}

type Messages []Message



