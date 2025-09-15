package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	dns := "host=localhost port=5432 user=postgres password=!Makar24052018 dname=social_network_db sslmode=disable"

	db, err := sqlx.Open("postgres", dns)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
