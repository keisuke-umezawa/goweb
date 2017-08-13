package route

import (
    "fmt"
    "testing"
    "net/http"
)

func TestNewRouter(t *testing.T) {

    name := "hoge"
    method := "GET"
    path := "/"
    routes := Routes{
        Route{
            name,
            method,
            path,
            hoge,
        },
    }

    router := NewRouter(routes)

    route := router.Get("hoge")
    
    actual := route.GetName()
    expected := name
    if expected != actual {
        t.Errorf("got %v\nwand %v",actual, expected)
    }
    
}

func hoge(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hoge\n")
}
