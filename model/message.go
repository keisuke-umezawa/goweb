package model

type Message struct {
    ID          uint        `json:"id" form:"id"`
    //User        User        `json:"user" form:"user"`
    UserID      uint        `json:"user_id" form:"user_id"`
    //Group       Group       `json:"group" form:"group"`
    GroupID     uint        `json:"group_id" form:"group_id"`
    Text        string      `json:"text" form:"text"`
    Mode        string      `json:"mode" form:"mode"`
}

type Messages []Message



