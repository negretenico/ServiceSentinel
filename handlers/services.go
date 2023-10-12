package handlers

import (
	app "ServiceSentinel/init"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetServices(c *gin.Context) {
	conf, err := app.Init()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, conf.Services)
}
