package model

type Message struct {
    ID          uint        `json:"id" form:"id"`
    SenderID    uint        `json:"sender_id" form:"sender_id"`
    GroupID     uint        `json:"group_id" form:"group_id"`
    Text        string      `json:"text" form:"text"`
    Mode        string      `json:"mode" form:"mode"`
}

type Messages []Message



