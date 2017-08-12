package main

import (
    "fmt"

    "github.com/keisuke-umezawa/goweb/model"
)

var currentId uint

var users model.Users

func init() {
    RepoCreateUser(model.User{Name: "keisuke"})
    RepoCreateUser(model.User{Name: "yusuke"})
}

func RepoFindUser(id uint) model.User {
    for _, u := range users {
        if u.Id == id {
            return u
        }
    }
    return model.User{}
}

// this is bad, I don't think it passes race conditions
func RepoCreateUser(u model.User) model.User {
    currentId += 1
    u.Id = currentId
    users = append(users, u)
    return u
}

func RepoDeleteUser(id uint) error {
    for i, u := range users {
        if u.Id == id {
            users = append(users[:i], users[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find User with id of %d to delete", id)
}
