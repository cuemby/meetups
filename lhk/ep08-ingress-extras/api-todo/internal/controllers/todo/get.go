package todo

import (
	"github.com/labstack/echo"
	"github.com/netsaj/episodio8-api-todo/internal/database"
	"github.com/netsaj/episodio8-api-todo/internal/database/models"
)

func Get(c echo.Context) error {
	userId := c.Get("user").(string)
	Id := c.Param("id")
	db := database.GetConnection()
	defer db.Close()
	var todo models.Todo
	if db.Model(&todo).First(&todo, "id = ? and user_id = ?", Id, userId).RecordNotFound() {
		return echo.ErrNotFound
	} else {
		return c.JSON(200, &todo)
	}
}
