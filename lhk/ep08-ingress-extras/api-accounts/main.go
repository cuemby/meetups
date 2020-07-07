package main

import (
	"github.com/labstack/echo"
	users "github.com/netsaj/episodio8-api-accounts/internal/controllers/users"
	"github.com/netsaj/episodio8-api-accounts/internal/database/migrations"
	"github.com/netsaj/episodio8-api-accounts/internal/middleware"
	cors_middleware "github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	//create database tables
	migrations.Migrate()
	// new echo server
	e := echo.New()
	e.Use(cors_middleware.CORSWithConfig(cors_middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	// add routes
	v1 := e.Group("/api/v1/accounts")
	// register new user
	v1.POST("/register", users.Register)
	v1.POST("/login", users.Login)
	user := v1.Group("/user", middleware.CustomJwtMiddleware())
	user.GET("/", users.GetLoggedUser)
	//start server
	e.Logger.Fatal(e.Start(":1323"))
}
