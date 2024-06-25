package main

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

const webPort = "80"

func main() {
	// connect to the db
	db := initDB()

	// create sessions
	session := initSession()

	// create channels

	// create wait groups
	wg := &sync.WaitGroup{}

	// setup the application config
	app := &Config{
		Session:  session,
		DB:       db,
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		Wait:     wg,
	}

	// set up mail

	// listen for web connections
	app.serve()
}

func (app *Config) serve() {
	// start http server
	srv := &http.Server{
		Addr:    ":" + webPort,
		Handler: app.Routes(),
	}

	app.InfoLog.Println("starting the server on port", webPort)

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}


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

func initSession() *scs.SessionManager {
	session := scs.New()
	session.Store = redisstore.New(initRedis())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}

func initRedis() *redis.Pool {
	redisPool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(os.Getenv("REDIS"))
		},
	}

	return redisPool

}
