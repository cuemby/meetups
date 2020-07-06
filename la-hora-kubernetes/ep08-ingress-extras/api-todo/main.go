package main

import (
	"github.com/labstack/echo"
	cors_middleware "github.com/labstack/echo/middleware"
	"github.com/netsaj/episodio8-api-todo/internal/controllers/todo"
	"github.com/netsaj/episodio8-api-todo/internal/database/migrations"
	"github.com/netsaj/episodio8-api-todo/internal/middleware"
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
	v1 := e.Group("/api/v1/todo")
	v1.Use(middleware.CustomJwtMiddleware)
	v1.GET("/", todo.GetList)
	v1.POST("/", todo.Create)
	v1.PUT("/", todo.Update)
	v1.DELETE("/:id", todo.Delete)
	v1.GET("/:id", todo.Get)
	//start server
	e.Logger.Fatal(e.Start(":1324"))
}
