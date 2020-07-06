package users

import (
	"github.com/labstack/echo"
	"github.com/netsaj/episodio8-api-accounts/internal/database"
	"github.com/netsaj/episodio8-api-accounts/internal/database/models"
	"github.com/netsaj/episodio8-api-accounts/internal/utils"
)

func Register(c echo.Context) error {
	var request models.User
	if err := c.Bind(&request); err != nil {
		return echo.ErrBadRequest
	}
	request.Password, _ = utils.HashPassword(request.Password)
	db := database.GetConnection()
	defer db.Close()
	if result := db.Save(&request); result.Error != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(201, map[string]interface{}{
		"user": &request,
	})
}
