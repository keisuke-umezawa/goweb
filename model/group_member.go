package model

type GroupMember struct {
    ID          uint        `json:"id" form:"id"`
    User        User
    UserID      uint
    Group       Group
    GroupID     uint
}
