package main

import (
    "log"
    "net/http"

    "github.com/keisuke-umezawa/goweb/db"
    "github.com/keisuke-umezawa/goweb/route"
    "github.com/keisuke-umezawa/goweb/model"
)

func main() {
    db, err := db.NewSqliteDB("db/database.db")
    if err != nil {
        log.Panic(err)
    }
    defer db.Close()

    // register model
    db.AutoMigrate(&model.User{})
    db.AutoMigrate(&model.Group{})
    db.AutoMigrate(&model.Message{})

    app := &Application{db: db}

    routes := route.Routes{
        route.Route{ "Index", "GET", "/", app.Index, },
        route.Route{ "UserIndex", "GET", "/users", app.UserIndex, },
        route.Route{ "UserShow", "GET", "/users/{id}", app.UserShow, },
        route.Route{ "UserCreate", "POST", "/users", app.UserCreate, },
        route.Route{ "GroupIndex", "GET", "/groups", app.GroupIndex, },
        route.Route{ "GroupShow", "GET", "/groups/{id}", app.GroupShow, },
        route.Route{ "GroupCreate", "POST", "/groups", app.GroupCreate, },
        route.Route{ "MessageIndex", "GET", "/messages", app.MessageIndex, },
        route.Route{ "MessagePost", "POST", "/messages", app.MessagePost, },
    }

    router := route.NewRouter(routes)

    // http://localhost:8080 でサービスを行う
    log.Fatal(http.ListenAndServe(":8080", router))
}
