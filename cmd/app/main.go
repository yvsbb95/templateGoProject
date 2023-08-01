package main

import (
	"log"
	"templateGoProject/pkg/config"
)

func main() {
	config := config.LoadConfig()
	log.Println("Config from local env", config)
}
