package modules

import (
	"encoding/json"
	"os"
)


type DataBaseConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
}

type Config struct {
	Database  DataBaseConfig   `json:"database"`
}

// читаем файл конфигурации
func ReadConfig(filePath string) (*Config, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var config *Config
	if err := json.Unmarshal(file,&config); err != nil {
		return nil,err
	}
	return config, nil 
}