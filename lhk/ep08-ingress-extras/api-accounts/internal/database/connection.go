/**
Connection to database with Gorm ORM
*/
package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"os"
)

// Db uri connection string
var DbUri string

// Setup connections params for postgres database
func init() {
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := "5432"
	DbUri = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)
}

// return a connection instance
// if all is okey return a DB connection instance, else call Panic(error)
func GetConnection() *gorm.DB {
	dbClient, err := gorm.Open(
		"postgres",
		DbUri)
	if err != nil {
		panic(fmt.Sprintf("failed connect to database: %s", err))
	}
	// Enable Logger, show detailed log
	dbClient.LogMode(true)
	return dbClient
}