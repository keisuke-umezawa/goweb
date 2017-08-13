package db

import (
    "testing"

    "github.com/keisuke-umezawa/goweb/model"
)

func TestNewSqliteDB(t *testing.T) {
    path := "test.db"
    db, err := NewSqliteDB(path)

    if err != nil {
        t.Errorf("got %v%v", err)
    }

    actual := db.HasTable(&model.User{})
    expected := true
    if expected != actual {
        t.Errorf("got %v\nwand %v",actual, expected)
    }
}
