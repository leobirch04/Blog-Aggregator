package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type JsonConfig struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFileName), nil
}

func Read() (JsonConfig, error) {
	name, err := getConfigFilePath()
	if err != nil {
		return JsonConfig{}, err
	}

	file, err := os.Open(name)
	if err != nil {
		return JsonConfig{}, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing file")
		}
	}(file)

	var config JsonConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return JsonConfig{}, err
	}

	return config, nil
}

func (con *JsonConfig) SetUser(userName string) error {
	con.CurrentUserName = userName
	return write(*con)
}

func write(con JsonConfig) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing file")
		}
	}(file)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(con)
	if err != nil {
		return err
	}
	return nil
}
