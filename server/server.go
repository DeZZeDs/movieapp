package server

import (
	"movieapp/database"
	"movieapp/logger"
	"net/http"
)

func ListenAndServe() {
	database.DB = database.ConnectorDB()

	logger.Trace("Server is listening...")
	logger.Critical("err: %v", http.ListenAndServe("localhost:5000", prepareRoutingTable()))
}
