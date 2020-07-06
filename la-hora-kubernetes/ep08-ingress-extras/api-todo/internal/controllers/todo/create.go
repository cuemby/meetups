package todo

import (
	"github.com/labstack/echo"
	"github.com/netsaj/episodio8-api-todo/internal/database"
	"github.com/netsaj/episodio8-api-todo/internal/database/models"
	uuid "github.com/satori/go.uuid"
)

func Create(c echo.Context) error {
	userId := c.Get("user").(string)
	var todo models.Todo
	c.Bind(&todo)
	db := database.GetConnection()
	defer db.Close()
	todo.UserId = uuid.FromStringOrNil(userId)
	if err := db.Save(&todo); err.Error != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(200, &todo)
}
