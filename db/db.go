package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    _ "github.com/jinzhu/gorm/dialects/postgres"
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

func NewPostgreDB(path string) (*DB, error) {
    db, err := gorm.Open("postgres", path)
    if err != nil {
        return nil, err
    }
    return &DB{db}, nil
}
