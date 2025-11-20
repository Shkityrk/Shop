package config

import "os"

// Config holds application configuration
type Config struct {
	GatewayPort        string
	AuthServiceURL     string
	ProductServiceURL  string
	CartServiceURL     string
}

// App is the global configuration instance
var App Config

func init() {
	App = Config{
		GatewayPort:        os.Getenv("GATEWAY_PORT"),
		AuthServiceURL:     os.Getenv("AuthServiceUrl"),
		ProductServiceURL:  os.Getenv("ProductServiceUrl"),
		CartServiceURL:     os.Getenv("Cart_service_url"),
	}
}

