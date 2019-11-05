package db

import (
	"context"
	"fmt"
	"github.com/jrkt/go-graphql-example/advanced/graphql/types"
	"math/rand"
	"time"
)

type User struct {
	id        string
	firstName string
	lastName  string
	email     *string
	createdAt types.DateTime
}

func (v *User) Id() string {
	return v.id
}

func (v *User) FirstName() string {
	return v.firstName
}

func (v *User) LastName() string {
	return v.lastName
}

func (v *User) FullName() string {
	return fmt.Sprintf("%s %s", v.firstName, v.lastName)
}

func (v *User) Email() *string {
	return v.email
}

func (v *User) CreatedAt() types.DateTime {
	return v.createdAt
}

func NewUser(firstName, lastName string, email *string) *User {
	return &User{
		id:        randomID(),
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		createdAt: types.NewDateTime(time.Now()),
	}
}

func FetchUserById(ctx context.Context, id string) (*User, error) {
	email := "jon.stevens@getweave.com"

	return &User{
		id:        id,
		firstName: "Jon",
		lastName:  "Stevens",
		email:     &email,
		createdAt: types.NewDateTime(time.Now().Add(-5 * time.Hour)),
	}, nil
}

func randomID() string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 16)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
