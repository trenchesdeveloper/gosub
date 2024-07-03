package main

import (
	"log"
	"sync"

	"github.com/alexedwards/scs/v2"
	"github.com/trenchesdeveloper/gosub/data"
	"github.com/upper/db/v4"
)

type Config struct {
	Session  *scs.SessionManager
	DB       db.Session
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Wait     *sync.WaitGroup
	Models   data.Models
}
