package main

import (
	"fmt"

	"github.com/fraser-isbester/pg-bq-go/internal/db"
	"github.com/fraser-isbester/pg-bq-go/internal/utils"
	_ "github.com/lib/pq"
)

func main() {

	// Load config file
	cfg, err := utils.LoadConfig("example/basic/config.yaml")
	if err != nil {
		fmt.Printf("Error reading YAML file: %v\n", err)
		return
	}

	// Connect & Test Source Database
	sourceDB, err := db.Open(cfg.DBs["source"].ConnectionString)
	if err != nil {
		fmt.Printf("Error connecting to source database: %v\n", err)
		return
	}
	defer sourceDB.Close()

	// Connect & Test Destination Database
	targetDB, err := db.Open(cfg.DBs["target"].ConnectionString)
	if err != nil {
		fmt.Printf("Error connecting to source database: %v\n", err)
		return
	}
	defer targetDB.Close()

	// Connect & Test Load balancer
	if cfg.LoadBalancer.Type == "pgcat" {
		fmt.Printf("insert pgcat logic here")
	} else {
		fmt.Printf("error: load balancer type not supported: %s\n", cfg.LoadBalancer.Type)
		return
	}

}
