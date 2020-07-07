package todo

import (
	"github.com/labstack/echo"
	"github.com/netsaj/episodio8-api-todo/internal/database"
	"github.com/netsaj/episodio8-api-todo/internal/database/models"
)

func GetList(c echo.Context) error {
	userId := c.Get("user")
	db := database.GetConnection()
	defer db.Close()
	var todos []models.Todo
	db.Model(&todos).Where("user_id = ?", &userId).Find(&todos)

	return c.JSON(200, map[string]interface{}{
		"todos": &todos,
	})
}
