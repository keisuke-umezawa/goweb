package model

import (
    "github.com/jinzhu/gorm"
)

type GroupMember struct {
    gorm.Model
    User        User
    UserID      uint
    Group       Group
    GroupID     uint
}
