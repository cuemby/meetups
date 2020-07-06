package users

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	. "github.com/netsaj/episodio8-api-accounts/internal/database"
	"github.com/netsaj/episodio8-api-accounts/internal/database/models"
	"net/http"
)

func GetLoggedUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	db := GetConnection()
	defer db.Close()

	var currentUser models.User
	if result := db.First(&currentUser, "id = ?", id); result.Error != nil {
		fmt.Print(result.Error.Error())
		return echo.ErrUnauthorized
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": currentUser,
	})
}
