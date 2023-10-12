package util

import (
	healthStatus "ServiceSentinel/models"
	"net/http"
)

func GetHealth(url string) healthStatus.HealthStatus {
	resp, err := http.Get(url)
	if err != nil {
		return healthStatus.UNKNOWN
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return healthStatus.UP
	}
	return healthStatus.DOWN
}
