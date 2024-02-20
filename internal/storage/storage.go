package storage

import (
	"database/sql"
	"fmt"
	"vanilla-server/internal/config"
)

func InitDB(cfg *config.Config) (*sql.DB, error) {
	fmt.Println("dbname", cfg.DBName)
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUser, cfg.DBPassword)
	fmt.Println(connectionString)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS weather (id SERIAL PRIMARY KEY, city TEXT, temperature REAL )")
	if err != nil {
		return nil, err
	}

	return db, nil

}
