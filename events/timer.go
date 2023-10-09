package events

import (
	app "ServiceSentinel/init"
	"fmt"
	"time"
)

func OnTime() {
	ticker := time.NewTicker(5 * time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			appConfig, err := app.Init()
			if err != nil {
				fmt.Println("Error loading in the config")
				return
			}
			fmt.Printf("Got the app config")
			fmt.Printf(appConfig.Services[0])
		}
	}
}
