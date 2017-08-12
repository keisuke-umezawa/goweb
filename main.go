package main

import (
    "log"
	"net/http"
)

func main() {
    router := NewRouter()

	// http://localhost:8080 でサービスを行う
	log.Fatal(http.ListenAndServe(":8080", router))
}
