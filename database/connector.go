package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"movieapp/config"
	"movieapp/logger"
)

var DB *sql.DB

func ConnectorDB() *sql.DB {
	path := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.User, config.Password, config.DatabaseName, config.Host, config.Port)

	dbConn, err := sql.Open("postgres", path)
	if err != nil {
		logger.Critical("Database connection failed: %v", err.Error())
	}
	if err = dbConn.Ping(); err != nil {
		logger.Critical("Database connection failed: %v", err.Error())
	}
	dbConn.SetMaxIdleConns(5)

	logger.Trace("Database connected")
	return dbConn
}
