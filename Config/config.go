package config

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	database *sql.DB
)

func init() {
	log.Println("Connection Pool Initializing")
	InitDB()
	log.Println("Connection Pool Initialiazed")
}

func InitDB() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "hydralisk123"
		dbname   = "ocbc-training"
	)

	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		// log.Fatal(err)
		fmt.Println("DB CONNECTION BROKEN")
	}
	database = db
}

func GetDB() *sql.DB {
	return database
}
