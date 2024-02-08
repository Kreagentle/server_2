package cmd

import (
	"log"
	"os"
)

func Logger() *log.Logger {
	return log.New(os.Stdout, "MyApp: ", log.Ldate|log.Ltime)
}
