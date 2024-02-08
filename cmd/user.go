package cmd

import (
	"sync"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var UsersDB = struct {
	sync.RWMutex
	Users map[int]User
}{
	Users: make(map[int]User),
}
