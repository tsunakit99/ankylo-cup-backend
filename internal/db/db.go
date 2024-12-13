package db

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq" // Postgres driver
	"github.com/tsunakit99/ankylo-cup-backend/internal/config"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		strconv.Itoa(cfg.DBPort), // int を文字列に変換
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Failed to connect DB: %v", err)
		return nil, err
	}

	// 確立を確認
	err = db.Ping()
	if err != nil {
		log.Printf("Failed to ping DB: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to DB")
	return db, nil
}
