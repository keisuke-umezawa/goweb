package main

import (
    "log"
	"net/http"

    "github.com/jinzhu/gorm"
    "github.com/keisuke-umezawa/goweb/db"
)

var database *gorm.DB

func main() {
    
    d := db.SqliteDatabase{Path: "db/database.db"}
    database = d.Connect()

    router := NewRouter()

	// http://localhost:8080 でサービスを行う
	log.Fatal(http.ListenAndServe(":8080", router))
}
