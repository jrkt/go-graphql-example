package db

import (
	"context"
	"github.com/jrkt/go-graphql-example/graphql/types"
	"math/rand"
	"time"
)

type User struct {
	id        string
	name      string
	email     *string
	createdAt types.DateTime
}

func (v *User) Id() string {
	return v.id
}

func (v *User) Name() string {
	return v.name
}

func (v *User) Email() *string {
	return v.email
}

func (v *User) CreatedAt() types.DateTime {
	return v.createdAt
}

func NewUser(name string, email *string) *User {
	return &User{
		id:        randomID(),
		name:      name,
		email:     email,
		createdAt: types.NewDateTime(time.Now()),
	}
}

func FetchUserById(ctx context.Context, id string) (*User, error) {
	email := "jon.stevens@getweave.com"

	return &User{
		id:        id,
		name:      "Jon Stevens",
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
