package events

import (
	app "ServiceSentinel/init"
	health "ServiceSentinel/util"
	"log"
	"strings"
	"time"
)

func OnTime() {
	ticker := time.NewTicker(5 * time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			conf, err := app.Init()
			if err != nil {
				return
			}
			hosts := []string{"chase", "teamfight", "git"}
			myServices := conf.Services
			for _, endpoint := range myServices {
				for _, host := range hosts {
					if strings.Contains(endpoint, host) {
						healthStatus := health.GetHealth(endpoint)
						log.Printf("The status for %s  is %s", endpoint, healthStatus)
					}
				}
			}
		}
	}
}
