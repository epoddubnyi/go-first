package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"my-first-go-project/config"
	"time"
)

var Db *sql.DB

func Connect() {
	var err error

	fmt.Println(config.Dsn)
	Db, err = sql.Open("pgx", config.Dsn)

	fmt.Println("TRY CONNECT", config.Dsn)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := Db.PingContext(ctx); err != nil {
		log.Fatal("DB connection error:", err)
	}
	log.Println("Connected to Postgres!")

}
