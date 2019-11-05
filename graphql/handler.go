package graphql

import (
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graph-gophers/graphql-transport-ws/graphqlws"
	"net/http"
)

func InitHandler() (http.Handler, error) {
	schema, err := Compile()
	if err != nil {
		return nil, nil
	}

	return graphqlws.NewHandlerFunc(schema, &relay.Handler{Schema: schema}), nil
}
