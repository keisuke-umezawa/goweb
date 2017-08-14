package model

type Group struct {
    ID          uint        `json:"id" form:"id"`
    Name        string      `json:"name" form:"name"`
}

type Groups []Group
