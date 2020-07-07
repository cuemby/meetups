package users

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/netsaj/episodio8-api-accounts/internal/database"
	"github.com/netsaj/episodio8-api-accounts/internal/database/models"
	"net/http"
	"time"
)

func Login(c echo.Context) (err error) {

	type Params struct {
		Email string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	u := new(Params)

	if err = c.Bind(u); err != nil {
		print(err)
		return echo.ErrBadRequest
	}
	DB := database.GetConnection()
	defer DB.Close()

	var user models.User
	if err := DB.Where("email = ?", u.Email).First(&user).Error; err != nil {
		print(err)
	}
	if user.CheckPassword(u.Password) {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)
		// Set custom claims
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = user.ID
		claims["Name"] = user.Name
		claims["Email"] = user.Email
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		// Create token with claims
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"token": t,
		})
	}
	return echo.ErrUnauthorized
}