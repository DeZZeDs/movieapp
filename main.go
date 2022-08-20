package main

import (
	"movieapp/database"
)

func main() {
	database.DB = database.ConnectorDB()
}
