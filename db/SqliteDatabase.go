package db

import (
	"os"
    "log"
	"path/filepath"
    "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
    "github.com/keisuke-umezawa/goweb/model"
)

type SqliteDatabase struct {
    Path string
}

func (self SqliteDatabase) Connect() *gorm.DB {
	dir := filepath.Dir(self.Path)
	db, err := gorm.Open("sqlite3", dir+"/database.db")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	if os.Getenv("AUTOMIGRATE") == "1" {
		db.AutoMigrate(
			&model.User{},
			&model.Message{},
			&model.Group{},
		)
	}

	return db
}
