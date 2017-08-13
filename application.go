package main

import (
    "log"

    "github.com/keisuke-umezawa/goweb/db"
)

type Application struct {
    db *db.DB
    logger *log.Logger
}
