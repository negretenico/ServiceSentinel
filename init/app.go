package init

import (
	config "ServiceSentinel/models"
	"encoding/json"
	"fmt"
	"os"
)

func Init() (config.AppConfig, error) {
	file, err := os.Open("config/local.json")
	if err != nil {
		fmt.Println("This is the wrong config file")
		return config.AppConfig{}, err // return zero inialized if not present
	}
	defer file.Close()
	var config config.AppConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		fmt.Println("Error reading in the file", err)
		return config.AppConfig{}, err // return zero inialized if not present
	}
	return config, nil
}
