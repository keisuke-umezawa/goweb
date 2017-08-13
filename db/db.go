package db

import (
    "os"

    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
    "github.com/keisuke-umezawa/goweb/model"
)

type DB struct {
    *gorm.DB
}

func NewSqliteDB(path string) (*DB, error) {
    db, err := gorm.Open("sqlite3", path)
    if err != nil {
        return nil, err
    }
    if os.Getenv("AUTOMIGATE") == "1" {
        db.AutoMigrate(
            &model.User{},
        )
    }
    if !db.HasTable(&model.User{}) {
        db.CreateTable(&model.User{})
    }
    // db.Create(&model.User{Name: "keisuke"})
    // db.Create(&model.User{Name: "yusuke"})
    return &DB{db}, nil
}
