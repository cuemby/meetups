package migrations

import (
	"github.com/netsaj/episodio8-api-todo/internal/database"
	"github.com/netsaj/episodio8-api-todo/internal/database/models"
)

func Migrate() {
	db := database.GetConnection()
	defer db.Close()
	db.AutoMigrate(
		&models.Todo{})
}
