package todo

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/netsaj/episodio8-api-todo/internal/database"
	"github.com/netsaj/episodio8-api-todo/internal/database/models"
	uuid "github.com/satori/go.uuid"
)

func Update(c echo.Context) error {
	userId := c.Get("user").(string)
	var todo models.Todo
	c.Bind(&todo)
	db := database.GetConnection()
	defer db.Close()
	todo.UserId = uuid.FromStringOrNil(userId)
	if db.Model(&todo).First(&models.Todo{}, "id = ? and user_id = ?", &todo.ID, userId).RecordNotFound() {
		return echo.ErrNotFound
	}
	todo.UserId = uuid.FromStringOrNil(userId)
	if err := db.Model(&todo).Save(&todo); err.Error != nil {
		fmt.Print(err.Error)
		return echo.ErrBadRequest
	} else {
		return c.JSON(200, &todo)
	}

}
