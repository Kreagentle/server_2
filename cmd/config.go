package cmd

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	Server struct {
		Port string `json:"port"`
		Host string `json:"host"`
	} `json:"server"`
}

func getEnvVars(filePath string) (Config, error) {
	var env Config
	f, err := ioutil.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return env, err
	}

	err = json.Unmarshal([]byte(f), &env)
	if err != nil {
		return env, err
	}
	return env, nil
}
