package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

type DB struct {
    *gorm.DB
}

func NewSqliteDB(path string) (*DB, error) {
    db, err := gorm.Open("sqlite3", path)
    if err != nil {
        return nil, err
    }
    return &DB{db}, nil
}
