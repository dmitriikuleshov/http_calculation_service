package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	conf "github.com/dmitriikuleshov/http_calculation_service/config"
	application "github.com/dmitriikuleshov/http_calculation_service/internal/application"
	calc "github.com/dmitriikuleshov/http_calculation_service/pkg/calculation"
)

func main() {
	// Load server configuration
	config, err := conf.LoadConfig(os.Getenv("CALC_CONFIG_PATH"))
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	host := config.Server.Host
	port := config.Server.Port

	// Create new handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		application.CalculateHandler(w, r, calc.Calc)
	})

	// Start server
	address := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("Server listening at %s\n", address)
	log.Fatal(http.ListenAndServe(address, handler))
}