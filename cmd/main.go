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

	// Configure Source Database for Logical Replication Publishing

	// Configure Target Database for Logical Replication Subscription

	// // Configure Replication object
	// replication := db.NewReplication(sourceDB, targetDB)

	// gets the database configuration of the Source DB
	sourceConf, err := db.GetConfig(sourceDB)
	if err != nil {
		fmt.Printf("Error getting database config: %v\n", err)
		return
	}
	fmt.Println(sourceConf)

	// gets the database configuration of the target DB
	targetConf, err := db.GetConfig(targetDB)
	if err != nil {
		fmt.Printf("Error getting database config: %v\n", err)
		return
	}
	fmt.Println(targetConf)

	// Assert replication configuration
	if sourceConf.WALLevel != "logical" {
		fmt.Printf("WAL Level must be set to 'logical' on the source database. Current value: %s)", sourceConf.WALLevel)
		return
	}
	if targetConf.HotStandby != "on" {
		fmt.Printf("Hot Standby must be enabled on the target database. Current value: %s)", targetConf.HotStandby)
		return
	}

	// Configure Source Publication

}
