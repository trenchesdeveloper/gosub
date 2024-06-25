package main

import (
	"log"
	"os"
	"time"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

const webPort = "80"

func main() {
	// connect to the db
	db := initDB()

	db.Ping()

	// create sessions

	// create channels

	// create wait groups

	// setup the application config

	// set up mail

	// listen for web connections
}

func initDB() db.Session {
	conn := connectToDB()

	if conn == nil {
		log.Fatal("could not connect to the database")
	}

	return conn
}

func connectToDB() db.Session {
	counts := 0

	dsn := os.Getenv("DSN")

	for {
		conn, err := openDB(dsn)

		if err != nil {
			log.Println("could not connect to the database", err)
			counts++

			if counts > 5 {
				return nil
			}

			time.Sleep(5 * time.Second)
		} else {
			log.Println("connected to the database")
			return conn
		}

	}
}

func openDB(dsn string) (db.Session, error) {
	url, err := postgresql.ParseURL(dsn)

	if err != nil {
		return nil, err
	}

	conn, err := postgresql.Open(url)

	if err != nil {
		return nil, err
	}

	err = conn.Ping()

	if err != nil {
		return nil, err
	}

	return conn, nil
}
