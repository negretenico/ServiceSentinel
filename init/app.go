package init

import (
	config "ServiceSentinel/models"
	"encoding/json"
	"fmt"
	"os"
)

func Init() (config.AppConfig, error) {
	file, err := os.Open("config/local.json")
	zero := config.AppConfig{}
	if err != nil {
		fmt.Println("This is the wrong config file")
		return zero, err // return zero inialized if not present
	}
	defer file.Close()
	var config config.AppConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		fmt.Println("Error reading in the file", err)
		return zero, err // return zero inialized if not present
	}
	return config, nil
}
