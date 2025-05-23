package main

import (
	"log"
	"system/cmd/api/routers"

	"system/internal/platform/migrate"

	"system/internal/auth/repo/postgre"
	"system/internal/platform/db"
)

func main() {

	database := db.ConnectDB()

	

	migrator := migrate.NewMigrator(database)
	if err := migrator.Run(); err != nil {
		log.Fatal("Migration failed:", err)
	}



	repo := postgre.NewUserRepository(database)

	routers.SetupRouter(repo)

	defer db.DisconnectDB(database)

}
