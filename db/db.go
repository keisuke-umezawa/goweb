package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "github.com/keisuke-umezawa/goweb/model"
)

type Datastore interface {
    insert(user model.User) model.User
}

type SqliteDB struct {
    *sql.DB
}

func NewSqliteDB(path string) (*SqliteDB, error) {
    db, err := sql.Open("sqlite3", path)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return &SqliteDB{db}, nil
}


func (db *SqliteDB) insert(user model.User) model.User {
    stmt, err := db.Prepare("INSERT INTO userinfo(name) values(keisuke)")
    checkError(err)
    res, err := stmt.Exec("astaxie", "研究開発部門", "2012-12-09")
    checkError(err)
    id, err := res.LastInsertId()
    checkError(err)

    user.Id = uint(id)
    return user
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
