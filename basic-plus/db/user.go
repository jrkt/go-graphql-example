package db

import (
	"context"
	"math/rand"
)

type User struct {
	id        string
	name      string
	email     *string
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

func NewUser(name string, email *string) *User {
	return &User{
		id:        randomID(),
		name:      name,
		email:     email,
	}
}

func FetchUserById(ctx context.Context, id string) (*User, error) {
	email := "jon.stevens@getweave.com"

	return &User{
		id:        id,
		name:      "Jon Stevens",
		email:     &email,
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
