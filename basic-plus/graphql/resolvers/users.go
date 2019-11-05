package resolvers

import (
	"context"
	"github.com/jrkt/go-graphql-example/basic-plus/db"
)

type userResolver struct{}

func NewUserResolver() *userResolver {
	return &userResolver{}
}

func (r *userResolver) GetUser(ctx context.Context, args struct{ Id string }) (*db.User, error) {
	return db.FetchUserById(ctx, args.Id)
}

func (r *userResolver) CreateUser(args struct {
	Name  string
	Email *string
}) (*db.User, error) {
	return db.NewUser(args.Name, args.Email), nil
}
