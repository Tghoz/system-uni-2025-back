package main

import (
	"system/internal/auth/models"
	"system/platform/db"
)

func main() {

	database := db.ConnectDB()
	database.AutoMigrate(&models.User{})

	defer db.DisconnectDB(database)

}
