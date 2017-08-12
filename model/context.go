package model

type Group struct {
    Id          uint        `gorm:"primary_key;Auto_INCREMENT" json:"id" form:"id"`
    Member_ids  []uint        `json:"context_id" form:"context_id"`
}
