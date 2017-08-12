package main

import (
    "encoding/json"
	"io"
	"io/ioutil"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!\n")
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(users); err != nil {
        panic(err)
    }
}

func UserShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "%s user index\n", vars["userId"])
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\nNot Found 404\n"))
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
    var user User
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &user); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    u := RepoCreateUser(user)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(u); err != nil {
        panic(err)
    }
}
