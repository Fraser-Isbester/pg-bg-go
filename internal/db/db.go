package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// type DB struct {
// 	*sql.DB
// 	conf *Config
// }

func Open(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging the database: %v", err)
	}

	return db, nil
}

type Config struct {
	WALLevel            string
	MaxWALSenders       int
	MaxReplicationSlots int
	WALSenderTimeout    string
	HotStandby          string
}

func GetWALLevel(db *sql.DB) (string, error) {
	var walLevel string
	err := db.QueryRow("SHOW wal_level").Scan(&walLevel)
	if err != nil {
		return "", fmt.Errorf("failed to get wal_level: %v", err)
	}
	return walLevel, nil
}

func GetMaxWALSenders(db *sql.DB) (int, error) {
	var maxWALSenders int
	err := db.QueryRow("SHOW max_wal_senders").Scan(&maxWALSenders)
	if err != nil {
		return 0, fmt.Errorf("failed to get max_wal_senders: %v", err)
	}
	return maxWALSenders, nil
}

func GetMaxReplicationSlots(db *sql.DB) (int, error) {
	var maxReplicationSlots int
	err := db.QueryRow("SHOW max_replication_slots").Scan(&maxReplicationSlots)
	if err != nil {
		return 0, fmt.Errorf("failed to get max_replication_slots: %v", err)
	}
	return maxReplicationSlots, nil
}

func GetWALSenderTimeout(db *sql.DB) (string, error) {
	var walSenderTimeout string
	err := db.QueryRow("SHOW wal_sender_timeout").Scan(&walSenderTimeout)
	if err != nil {
		return "", fmt.Errorf("failed to get wal_sender_timeout: %v", err)
	}
	return walSenderTimeout, nil
}

func GetHotStandby(db *sql.DB) (string, error) {
	var hotStandby string
	err := db.QueryRow("SHOW hot_standby").Scan(&hotStandby)
	if err != nil {
		return "", fmt.Errorf("failed to get hot_standby: %v", err)
	}
	return hotStandby, nil
}

func GetConfig(db *sql.DB) (*Config, error) {
	walLevel, err := GetWALLevel(db)
	if err != nil {
		return nil, err
	}

	maxWALSenders, err := GetMaxWALSenders(db)
	if err != nil {
		return nil, err
	}

	maxReplicationSlots, err := GetMaxReplicationSlots(db)
	if err != nil {
		return nil, err
	}

	walSenderTimeout, err := GetWALSenderTimeout(db)
	if err != nil {
		return nil, err
	}

	GetHotStandby, err := GetHotStandby(db)
	if err != nil {
		return nil, err
	}

	return &Config{
		WALLevel:            walLevel,
		MaxWALSenders:       maxWALSenders,
		MaxReplicationSlots: maxReplicationSlots,
		WALSenderTimeout:    walSenderTimeout,
		HotStandby:          GetHotStandby,
	}, nil
}
