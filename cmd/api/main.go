package main

import (
	"system/cmd/api/routers"
	"system/internal/auth/model"
	"system/internal/auth/repo/postgre"
	"system/internal/platform/db"
)

func main() {

	database := db.ConnectDB()
	database.AutoMigrate(&model.User{})

	repo := postgre.NewUserRepository(database)

	routers.SetupRouter(repo)

	defer db.DisconnectDB(database)

}
