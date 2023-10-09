package config

type ServerConfig struct {
	Port int `json:"port"`
}
type AppConfig struct {
	Server   ServerConfig `json:"sever"`
	Services []string     `json:"services"`
}
