package main

import (
	"github.com/Kreagentle/server_2/cmd"
)

func main() {
	logger := cmd.Logger()

	// load environment variables from `env.json`.
	envVars, err := cmd.GetEnvVars("env.json")
	if err != nil {
		logger.Printf("Error reading config file: %s \n", err.Error)
		return
	}

	// run program.
	server := cmd.Server()

	// start UI process
	logger.Printf("The webserver is started at the port: %s \n")
	if err := server.Serve(s); err != nil {
		logger.Printf("Couldnt start webserver: %s \n", err.Error)
		return
	}
}
