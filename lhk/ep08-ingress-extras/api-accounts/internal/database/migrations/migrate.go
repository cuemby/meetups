package migrations

import (
	"github.com/netsaj/episodio8-api-accounts/internal/database"
	"github.com/netsaj/episodio8-api-accounts/internal/database/models"
)

func Migrate() {
	db := database.GetConnection()
	defer db.Close()
	db.AutoMigrate(
		&models.User{})

	db.Model(&models.User{}).AddUniqueIndex("email_uk", "email")
}
