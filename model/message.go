package model

type Message struct {
    Id          uint        `gorm:"primary_key;Auto_INCREMENT" json:"id" form:"id"`
    SenderId    uint        `json:"sender_id" form:"sender_id"`
    GroupId     uint        `json:"group_id" form:"group_id"`
    Text        string      `json:"text" form:"text"`
    Mode        string      `json:"mode" form:"mode"`
}

type Messages []Message



