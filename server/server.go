package server

import (
	"movieapp/config"
	"movieapp/database"
	"movieapp/logger"
	"net/http"
)

func Start() {
	config.SetConfig()
	database.DB = database.ConnectorDB()

	logger.Trace("Server is listening...")
	logger.Critical("err: %v", http.ListenAndServe("localhost:5000", prepareRoutingTable()))
}
