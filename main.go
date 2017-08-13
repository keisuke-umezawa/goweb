package main

import (
    "log"
    "net/http"

    "github.com/keisuke-umezawa/goweb/db"
    "github.com/keisuke-umezawa/goweb/route"
)

func main() {
    db, err := db.NewSqliteDB("db/database.db")
    if err != nil {
        log.Panic(err)
    }
    app := &Application{db: db}

    routes := route.Routes{
        route.Route{
            "Index",
            "GET",
            "/",
            app.Index,
        },
        route.Route{
            "UserIndex",
            "GET",
            "/users",
            app.UserIndex,
        },
        route.Route{
            "UserShow",
            "GET",
            "/users/{userId}",
            app.UserShow,
        },
        route.Route{
            "UserCrate",
            "POST",
            "/users",
            app.UserCreate,
        },
    }

    router := route.NewRouter(routes)

    // http://localhost:8080 でサービスを行う
    log.Fatal(http.ListenAndServe(":8080", router))
}
