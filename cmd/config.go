package cmd

import (
	"encoding/json"
	"file"
)

func getEnvVars(filePath string) (EnvVars, error) {
	var env EnvVars
	f, err := file.ReadFile(filePath)
	if err != nil {
		return env, err
	}

	err = json.Unmarshal([]byte(f), &env)
	if err != nil {
		return env, err
	}
	return env, nil
}
