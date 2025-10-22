package config

import (
	_ "log"
	"os"
)

type Config struct {
	Gateway_port        string
	Auth_service_url    string
	Product_service_url string
	Cart_service_url    string
}

var App Config

func init() {
	App = Config{
		Gateway_port:        os.Getenv("GATEWAY_PORT"),
		Auth_service_url:    os.Getenv("AuthServiceUrl"),
		Product_service_url: os.Getenv("ProductServiceUrl"),
		Cart_service_url:    os.Getenv("Cart_service_url"),
	}
}
