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
		config.Database.DbUser, config.Database.DbPassword, config.Database.DbName, config.Database.DbHost, config.Database.DbPort)

	dbConn, err := sql.Open(config.Database.DriverName, path)
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
