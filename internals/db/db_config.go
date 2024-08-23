package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

//var DATABASE_URI string = "postgres:postgres@tcp(localhost:55000)/e-learning"

func Connect(dbHost, dbUser, dbPassword string, dbPort int) error {
	var err error

	var DSN string = fmt.Sprintf("host=%s user=%s password=%s dbname=e-learning port=%d sslmode=disable", dbHost, dbUser, dbPassword, dbPort)

	Database, err = gorm.Open(postgres.Open(DSN), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	return nil
}
