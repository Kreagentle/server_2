package main

import (
	"net/http"
	"time"

	"github.com/Kreagentle/server_2/cmd"
	"github.com/gobuffalo/buffalo/servers"
)

func main() {
	logger := cmd.Logger()

	// load environment variables from `env.json`.
	envVars, err := cmd.GetEnvVars("env.json")
	if err != nil {
		logger.Printf("Error reading config file: %s \n", err.Error)
		return
	}

	var s servers.Server
	serverconf := &http.Server{
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  90 * time.Second,
		WriteTimeout: 900 * time.Second,
		Addr:         ":" + envVars.Server.Port,
	}
	s = servers.Wrap(serverconf)

	// run program.
	server := cmd.Server()

	// start UI process
	logger.Printf("The webserver is started at the port: %s \n")
	if err := server.Serve(s); err != nil {
		logger.Printf("Couldnt start webserver: %s \n", err.Error)
		return
	}
}
