package main

import (
    "database/sql"
    "log"
)

type Application struct {
    db *sql.DB
    logger *log.Logger
}
