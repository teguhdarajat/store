package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	err = db.Ping()
	if err != nil {
		log.Printf("Ping failed. Check if your IP is allowed in Supabase Dashboard -> Network Restrictions")
		return nil, err
	}

	log.Println("Database connected successfully")
	return db, nil
}
