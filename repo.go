package main

import "fmt"

var currentId int

var users Users

func init() {
    RepoCreateUser(User{Name: "keisuke"})
    RepoCreateUser(User{Name: "yusuke"})
}

func RepoFindUser(id int) User {
    for _, u := range users {
        if u.Id == id {
            return u
        }
    }
    return User{}
}

// this is bad, I don't think it passes race conditions
func RepoCreateUser(u User) User {
    currentId += 1
    u.Id = currentId
    users = append(users, u)
    return u
}

func RepoDeleteUser(id int) error {
    for i, u := range users {
        if u.Id == id {
            users = append(users[:i], users[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find User with id of %d to delete", id)
}
