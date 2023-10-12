package models

type HealthStatus string

const (
	UP      HealthStatus = "UP"
	DOWN    HealthStatus = "DOWN"
	UNKNOWN HealthStatus = "UNKNOWN"
)
