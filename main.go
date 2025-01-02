package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/maathias/capstone/api"
	"github.com/maathias/capstone/db"
	"github.com/maathias/capstone/edge"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Envrironment
	service := os.Getenv("SERVICE")
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPass := os.Getenv("REDIS_PASS")

	if service == "" || redisAddr == "" {
		fmt.Println("Missing environment variables")
		os.Exit(1)
	}

	if service != "edge" && service != "api" {
		fmt.Println("Invalid service")
		os.Exit(1)
	}

	// Prometheus
	go func(){
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)

	}()

	// Redis
	db.Connect(service, redisAddr, redisPass)

	// Run service
	if service == "edge" {
		edge.Run()
	} else if service == "api" {
		api.Run()
	} else {
		fmt.Println("Invalid service")
		os.Exit(1)
	}

	select {}
}
