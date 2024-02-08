package cmd

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
)

func getUsersHandler(c buffalo.Context) error {
	UsersDB.RLock()
	defer UsersDB.RUnlock()

	users := make([]User, 0, len(UsersDB.Users))
	for _, user := range UsersDB.Users {
		users = append(users, user)
	}

	return c.Render(http.StatusOK, render.JSON(users))
}

func createUserHandler(c buffalo.Context) error {
	var newUser User
	err := c.Bind(&newUser)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}

	UsersDB.Lock()
	defer UsersDB.Unlock()

	newUser.ID = len(UsersDB.Users) + 1
	UsersDB.Users[newUser.ID] = newUser

	return c.Render(http.StatusOK, render.JSON(newUser))
}

func getUserHandler(c buffalo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}

	UsersDB.RLock()
	defer UsersDB.RUnlock()

	user, ok := UsersDB.Users[id]
	if !ok {
		return c.Error(http.StatusNotFound, fmt.Errorf("User not found"))
	}

	return c.Render(http.StatusOK, render.JSON(user))
}

func updateUserHandler(c buffalo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}

	UsersDB.Lock()
	defer UsersDB.Unlock()

	_, ok := UsersDB.Users[id]
	if !ok {
		return c.Error(http.StatusNotFound, fmt.Errorf("User not found"))
	}

	var updatedUser User
	err = c.Bind(&updatedUser)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}

	UsersDB.Users[id] = updatedUser

	return c.Render(http.StatusOK, render.JSON(updatedUser))
}

func deleteUserHandler(c buffalo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}

	UsersDB.Lock()
	defer UsersDB.Unlock()

	_, ok := UsersDB.Users[id]
	if !ok {
		return c.Error(http.StatusNotFound, fmt.Errorf("User not found"))
	}

	delete(UsersDB.Users, id)

	return c.Render(http.StatusNoContent, nil)
}
