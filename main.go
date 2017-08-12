package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// 単純なハンドラ
	r.HandleFunc("/", Index)

	// パスに変数を埋め込み
	r.HandleFunc("/users", UserIndex)

	// パスに変数を埋め込み
	r.HandleFunc("/users/{userId}", UserIndexId)

	// パス変数で正規表現を使用

	// マッチするパスがない場合のハンドラ
	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	// http://localhost:8080 でサービスを行う
	http.ListenAndServe(":8080", r)
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index!\n"))
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UserIndex!\n"))
}

func UserIndexId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "%s user index\n", vars["userId"])
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\nNot Found 404\n"))
}
