package main

import (
	"log"
	"system/cmd/api/routers"

	"system/internal/platform/migrate"

	"system/internal/platform/db"
	"system/internal/repo"
)

func main() {

	database := db.ConnectDB()

	migrator := migrate.NewMigrator(database)
	if err := migrator.Run(); err != nil {
		log.Fatal("Migration failed:", err)
	}

	cont := repo.NewRepositoryContainer(database)

	routers.SetupRouter(cont)

	defer db.DisconnectDB(database)

}
