package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bohexists/task-manager-svc/config"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectToDB(cfg config.Config) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	log.Println("Successfully connected to database")
}
