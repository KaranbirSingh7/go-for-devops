package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DBUsername = "postgres"
	DBPassword = "postgres"
	DBHost     = "localhost"
	DBPort     = 5432
	Database   = "mydb"
)

func main() {
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", DBUsername, DBPassword, DBHost, DBPort, Database)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// create connection
	conn, err := pgxpool.New(ctx, dbURI)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	log.Println("DB connection successful.")
}
