package config

import (
	"encoding/json"
	"io/ioutil"
	"movieapp/logger"
)

type DbConfig struct {
	DriverName string `json:"driver_name"`
	DbUser     string `json:"db_user"`
	DbPassword string `json:"db_password"`
	DbName     string `json:"db_name"`
	DbHost     string `json:"db_host"`
	DbPort     string `json:"db_port"`
}

var Database DbConfig

func SetConfig() {
	file, err := ioutil.ReadFile("config/db_config.json")
	if err != nil {
		logger.Critical("SetConfig - Error ReadFile: %s\", err.Error()")
	}

	err = json.Unmarshal(file, &Database)
	if err != nil {
		logger.Critical("SetConfig - Error Unmarshaling file: %s\", err.Error()")
	}
}
