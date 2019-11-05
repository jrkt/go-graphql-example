package graphql

import (
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr"
	"github.com/graph-gophers/graphql-go"
	"github.com/jrkt/go-graphql-example/advanced/graphql/resolvers"
	"strings"
)

var (
	SchemaBox packr.Box
)

func init() {
	SchemaBox = packr.NewBox("./schema")
}

func Compile() (*graphql.Schema, error) {
	var s string

	err := SchemaBox.Walk(func(path string, file packd.File) error {
		if strings.HasSuffix(path, ".graphql") {
			fileString, err := SchemaBox.FindString(path)
			if err != nil {
				return err
			}

			s += `
` + fileString
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return graphql.ParseSchema(s, resolvers.Init())
}
