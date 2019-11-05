package resolvers

import (
	"context"
	"github.com/jrkt/go-graphql-example/advanced/db"
	"math/rand"
	"time"
)

type userResolver struct {
	userEvents     chan *db.User
	userSubscriber chan *userSubscriber
}

func NewUserResolver() *userResolver {
	r := &userResolver{
		userEvents:     make(chan *db.User),
		userSubscriber: make(chan *userSubscriber),
	}

	go r.startNotificationSubscription()

	return r
}

func (r *userResolver) GetUser(ctx context.Context, args struct{ Id string }) (*db.User, error) {
	return db.FetchUserById(ctx, args.Id)
}

func (r *userResolver) CreateUser(args struct {
	FirstName  string
	LastName  string
	Email *string
}) *db.User {
	u := db.NewUser(args.FirstName, args.LastName, args.Email)

	go func() {
		select {
		case r.userEvents <- u:
		case <-time.After(1 * time.Second):
		}
	}()

	return u
}

type userSubscriber struct {
	stop   <-chan struct{}
	events chan<- *db.User
}

func (r *userResolver) startNotificationSubscription() {
	subscribers := map[string]*userSubscriber{}
	unsubscribe := make(chan string)

	for {
		select {
		case id := <-unsubscribe:
			delete(subscribers, id)
		case s := <-r.userSubscriber:
			subscribers[randomID()] = s
		case e := <-r.userEvents:
			for id, s := range subscribers {
				go func(id string, s *userSubscriber) {
					select {
					case <-s.stop:
						unsubscribe <- id
						return
					default:
					}

					select {
					case <-s.stop:
						unsubscribe <- id
					case s.events <- e:
					case <-time.After(time.Second):
					}
				}(id, s)
			}
		}
	}
}

func (r *userResolver) Users(ctx context.Context) <-chan *db.User {
	c := make(chan *db.User)
	r.userSubscriber <- &userSubscriber{events: c, stop: ctx.Done()}

	return c
}

func randomID() string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 16)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
