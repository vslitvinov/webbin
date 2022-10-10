package web

import (
	"encoding/json"
	"os"
)

// читаем файл конфигурации
func ReadData(data *interface{}, filePath string) (interface{}, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	
	if err := json.Unmarshal(file,&data); err != nil {
		return nil,err
	}
	return data, nil 
}