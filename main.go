package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/n17ali/bank/api"
	db "github.com/n17ali/bank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://postgres:secret@localhost:5432/bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
