package main

import (
    "encoding/json"
    "io"
    "io/ioutil"
    "net/http"
    "strconv"
    "os/exec"

    "github.com/gorilla/mux"
    "github.com/keisuke-umezawa/goweb/model"
)

func (app *Application) Index(w http.ResponseWriter, r *http.Request) {
    if out, err := exec.Command("./script/hello").Output(); err != nil {
        w.Header().Set("Content-Type", "text; charset=UTF-8")
        w.WriteHeader(500) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
        return
    } else {
        if err := json.NewEncoder(w).Encode(string(out)); err != nil {
            panic(err)
        }
    }
}

func (app *Application) UserIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    users := model.Users{}
    app.db.Debug().Find(&users)
    if err := json.NewEncoder(w).Encode(users); err != nil {
        panic(err)
    }
}

func (app *Application) UserShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"]);
    if err != nil {
        panic(err)
    }
    var user model.User
    app.db.Debug().First(&user, id)
    if user.ID > uint(0) {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(user); err != nil {
            panic(err)
        }
        return
    }

    // if we do not find it. 404.
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusNotFound)
    if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
        panic(err)
    }
}

func (app *Application) UserCreate(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }

    // json -> User
    var user model.User
    if err := json.Unmarshal(body, &user); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
        return
    }

    // register into DB
    if err := app.db.Debug().Create(&user); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
        return
    }

    // make response
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(user); err != nil {
        panic(err)
    }
}

func (app *Application) GroupIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    groups := model.Groups{}
    app.db.Debug().Find(&groups)
    if err := json.NewEncoder(w).Encode(groups); err != nil {
        panic(err)
    }
}

func (app *Application) GroupShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"]);
    if err != nil {
        panic(err)
    }
    var group model.Group
    app.db.Debug().First(&group, id)
    if group.ID > uint(0) {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(group); err != nil {
            panic(err)
        }
        return
    }

    // if we do not find it. 404.
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusNotFound)
    if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
        panic(err)
    }
}

func (app *Application) GroupCreate(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }

    // json -> Group
    var group model.Group
    if err := json.Unmarshal(body, &group); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
        return
    }

    // register into DB
    if err := app.db.Debug().Create(&group); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
        return
    }

    // make response
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(group); err != nil {
        panic(err)
    }
}

func (app *Application) MessageIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    messages := model.Messages{}
    app.db.Debug().Find(&messages)
    if err := json.NewEncoder(w).Encode(messages); err != nil {
        panic(err)
    }
}

func (app *Application) MessagePost(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }

    // json -> Message
    var message model.Message
    if err := json.Unmarshal(body, &message); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
        return
    }
    // register into DB
        if err := app.db.Debug().Create(&message); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
        return
    }

    // make response
    response := model.Message{
        UserID: 1,
        GroupID: message.GroupID,
        Text: message.Text + "!!!", Mode: "dialogue"}

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        panic(err)
    }
    app.db.Debug().Create(&response)
}
