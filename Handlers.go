package main

import (
    "encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!\n")
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
    users := Users{
        User{Id: "00000001", Name: "keisuke"},
        User{Id: "00000002", Name: "yusuke"},
    }
    json.NewEncoder(w).Encode(users)
}

func UserShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "%s user index\n", vars["userId"])
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\nNot Found 404\n"))
}
