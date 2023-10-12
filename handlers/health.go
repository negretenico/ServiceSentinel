package handlers

import (
	app "ServiceSentinel/init"
	health "ServiceSentinel/util"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {
	urlFromPathVariable := c.Param("url")
	conf, err := app.Init()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error reading config": err})
		return
	}
	myServices := conf.Services
	url := "nil"
	for _, endpoint := range myServices {
		if strings.Contains(endpoint, urlFromPathVariable) {
			url = endpoint
			break
		}
	}
	if url == "nil" {
		c.JSON(http.StatusBadRequest, gin.H{"Could not find url for ": urlFromPathVariable})
		return
	}
	healthStatus := health.GetHealth(url)
	c.JSON(http.StatusOK, gin.H{fmt.Sprintf("Endpoint:%s", urlFromPathVariable): healthStatus})
}
