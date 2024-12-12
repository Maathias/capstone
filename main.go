package main

import (
	"fmt"
	"os"

	"github.com/maathias/capstone/api"
	"github.com/maathias/capstone/db"
	"github.com/maathias/capstone/edge"
)

func main() {
	service := os.Getenv("SERVICE")
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPass := os.Getenv("REDIS_PASS")

	if service == "" || redisAddr == "" || redisPass == "" {
		fmt.Println("Missing environment variables")
		os.Exit(1)
	}

	if service != "edge" && service != "api" {
		fmt.Println("Invalid service")
		os.Exit(1)
	}

	db.Connect(service, redisAddr, redisPass)

	if service == "edge" {
		edge.Run()
	} else if service == "api" {
		api.Run()
	} else {
		fmt.Println("Invalid service")
		os.Exit(1)
	}
}
