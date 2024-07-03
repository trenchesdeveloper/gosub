package data

import (
	"time"

	"github.com/upper/db/v4"
)

const dbTimeout = time.Second * 3

var dbInstance db.Session

// New is the function used to create an instance of the data package. It returns the type
// Model, which embeds all the types we want to be available to our application.
func New(dbPool db.Session) Models {
	dbInstance = dbPool

	return Models{
		User: User{},
		Plan: Plan{},
	}
}

// Models is the type for this package. Note that any model that is included as a member
// in this type is available to us throughout the application, anywhere that the
// app variable is used, provided that the model is also added in the New function.
type Models struct {
	User User
	Plan Plan
}
