package main

import (
    "log"
	"net/http"

    "github.com/jinzhu/gorm"
    "github.com/keisuke-umezawa/goweb/db"
)

var database *gorm.DB

func main() {
    db, err := db.NewSqliteDB("db/database.db")
    if err != nil {
        log.Panic(err)
    }
    app := &Application{db: db}

    routes := Routes{
        Route{
            "Index",
            "GET",
            "/",
            app.Index,
        },
        Route{
            "UserIndex",
            "GET",
            "/users",
            app.UserIndex,
        },
        Route{
            "UserShow",
            "GET",
            "/users/{userId}",
            app.UserShow,
        },
        Route{
            "UserCrate",
            "POST",
            "/users",
            app.UserCreate,
        },
    }

    router := NewRouter(routes)

	// http://localhost:8080 でサービスを行う
	log.Fatal(http.ListenAndServe(":8080", router))
}
