package cmd

import (
	"github.com/gobuffalo/buffalo"
)

func Server() *buffalo.App {
	server := buffalo.New(buffalo.Options{})

	server.GET("/users", getUsersHandler)
	server.POST("/users", createUserHandler)
	server.GET("/users/{id:[0-9]+}", getUserHandler)
	server.PUT("/users/{id:[0-9]+}", updateUserHandler)
	server.DELETE("/users/{id:[0-9]+}", deleteUserHandler)

	return server
}
