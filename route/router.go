package route

import (
    "net/http"

    "github.com/gorilla/mux"
)


func NewRouter(routes Routes) *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }

    router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
    return router
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla!\nNot Found 404\n"))
}
