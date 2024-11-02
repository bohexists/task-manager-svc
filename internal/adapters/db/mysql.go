package db

import (
	"database/sql"
	"fmt"
	"github.com/bohexists/task-manager-svc/internal/config"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectToDB(cfg config.Config) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	maxAttempts := 10
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		DB, err = sql.Open("mysql", dsn)
		if err == nil && DB.Ping() == nil {
			log.Println("Successfully connected to database")
			return
		}

		log.Printf("Attempt %d/%d: failed to connect to database: %v", attempts, maxAttempts, err)

		time.Sleep(5 * time.Second)
	}

	log.Fatalf("Failed to connect to database after %d attempts", maxAttempts)
}
