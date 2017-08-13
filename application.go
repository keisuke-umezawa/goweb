package main

import (
    "log"

    "github.com/keisuke-umezawa/goweb/db"
)

type Application struct {
    db db.Datastore
    logger *log.Logger
}
