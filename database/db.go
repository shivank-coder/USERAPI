package database

import (
	"context"
	"fmt"
	"go-practice-app/config"
	"log"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func Connect() {
	dbHost := config.GetEnv("DB_HOST")
	dbUser := config.GetEnv("DB_USER")
	dbPassword := config.GetEnv("DB_PASSWORD")
	dbName := config.GetEnv("DB_NAME")
	dbPort := config.GetEnv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	var err error
	DB, err = pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	log.Println("Database connected successfully")
}
