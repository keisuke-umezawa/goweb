package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/context"
)

func main() {
	r := mux.NewRouter()

	// 単純なハンドラ
	r.HandleFunc("/", YourHandler)

	// パスに変数を埋め込み
	r.HandleFunc("/hello/{name}", VarsHandler)

	// パス変数で正規表現を使用
	r.HandleFunc("/hello/{name}/{age:[0-9]+}", RegexHandler)

	// クエリ文字列の取得
	r.HandleFunc("/hi/", QueryStringHandler)

	// 静的ファイルの提供
	// $PROROOT/assets/about.html が http://localhost:8080/assets/about.html でアクセスできる
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	// リダイレクト
	r.HandleFunc("/moved", RedirectHandler)

	// マッチするパスがない場合のハンドラ
	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	// 複数のハンドラで共通の処理を実行する
	// 今回はcontextのセットとゲットを試しているが、同じパターンでDBの初期化や認証処理やログ書き出しなどにも応用できる
	// ハンドラを引き渡すには http.Handler 型を使い func(http.ResponseWriter, *http.Request) から http.Handler への変換には http.HandlerFunc を利用する
	// http.Handler をハンドラとして登録する場合は Router.Handle を利用する
	r.Handle("/some1", UseContext(http.HandlerFunc(SomeHandler1)))
	r.Handle("/some2", UseContext(http.HandlerFunc(SomeHandler2)))

	// http://localhost:8080 でサービスを行う
	http.ListenAndServe(":8080", r)
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func VarsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "%s Loves Gorilla\n", vars["name"])
}

func RegexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "%s is %s years old\n", vars["name"], vars["age"])
}

func QueryStringHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	fmt.Fprintf(w, "%s Loves Gorilla\n", q.Get("name"))
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusFound)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\nNot Found\n"))
}

const MyContextKey = 1

func UseContext(handler http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, MyContextKey, "Call SomeMiddleware")
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func SomeHandler1(w http.ResponseWriter, r *http.Request) {
	contextVal := context.Get(r, MyContextKey)
	fmt.Fprintf(w, "%s Call SomeHandler1", contextVal)
}

func SomeHandler2(w http.ResponseWriter, r *http.Request) {
	contextVal := context.Get(r, MyContextKey)
	fmt.Fprintf(w, "%s Call SomeHandler2", contextVal)
}
