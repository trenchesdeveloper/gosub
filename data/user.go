package data

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/upper/db/v4"
	"golang.org/x/crypto/bcrypt"
)

// User is the structure which holds one user from the database.
type User struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	Password  string
	Active    int
	IsAdmin   int
	CreatedAt time.Time
	UpdatedAt time.Time
	Plan      *Plan
}

// GetAll returns a slice of all users, sorted by last name
func (u *User) GetAll() ([]*User, error) {
	_, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var users []*User

	userCollection := dbInstance.Collection("users")

	err := userCollection.Find().OrderBy("last_name").All(&users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetByEmail returns one user by email
func (u *User) GetByEmail(email string) (*User, error) {
	_, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var user User

	userCollection := dbInstance.Collection("users")

	err := userCollection.Find(db.Cond{"email": email}).One(&user)

	if err != nil {
		return nil, err
	}

	// use query builder
	q := dbInstance.SQL().
		Select("p.id", "p.plan_name", "p.plan_amount", "p.created_at", "p.updated_at").
		From("plans p").
		LeftJoin("user_plans up", "p.id = up.plan_id").
		Where("up.user_id = $1", user.ID)

	var plan Plan

	if err := q.One(&plan); err == nil {
		user.Plan = &plan
	}

	return &user, nil
}

// GetOne returns one user by id
func (u *User) GetOne(id int) (*User, error) {
	_, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	userCollection := dbInstance.Collection("users")

	var user User

	err := userCollection.Find(db.Cond{"id": id}).One(&user)


	if err != nil {
		return nil, err
	}

	q := dbInstance.SQL().
		Select("p.id", "p.plan_name", "p.plan_amount", "p.created_at", "p.updated_at").
		From("plans p").
		LeftJoin("user_plans up", "p.id = up.plan_id").
		Where("up.user_id = $1", user.ID)

	var plan Plan

	if err := q.One(&plan); err == nil {
		user.Plan = &plan
	}

	if err == nil {
		user.Plan = &plan
	} else {
		log.Println("Error getting plan", err)
	}

	return &user, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *User) Update() error {
	_, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	userCollection := dbInstance.Collection("users")

	err := userCollection.Find(db.Cond{"id": u.ID}).Update(
		&User{
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Active:    u.Active,
			UpdatedAt: time.Now(),

		},
	)

	if err != nil {
		return err
	}

	return nil
}

// Delete deletes one user from the database, by User.ID
func (u *User) Delete() error {
	_, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	userCollection := dbInstance.Collection("users")

	err := userCollection.Find(db.Cond{"id": u.ID}).Delete()

	if err != nil {
		return err
	}

	return nil
}

// DeleteByID deletes one user from the database, by ID
func (u *User) DeleteByID(id int) error {
	_, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	userCollection := dbInstance.Collection("users")

	err := userCollection.Find(db.Cond{"id": id}).Delete()

	if err != nil {
		return err
	}

	return nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (u *User) Insert(user User) (int, error) {
	_, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return 0, err
	}

	var newID int

	userCollection := dbInstance.Collection("users")

	insertResult, err := userCollection.Insert(map[string]interface{}{
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"password":   hashedPassword,
		"user_active": user.Active,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	})

	if err != nil {
		return 0, err
	}

	newID = insertResult.ID().(int)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// ResetPassword is the method we will use to change a user's password.
func (u *User) ResetPassword(password string) error {
	_, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	userCollection := dbInstance.Collection("users")

	err = userCollection.Find(db.Cond{"id": u.ID}).Update(map[string]interface{}{
		"password": hashedPassword,
	})

	if err != nil {
		return err
	}

	return nil
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
// with the hash we have stored for a given user in the database. If the password
// and hash match, we return true; otherwise, we return false.
func (u *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
