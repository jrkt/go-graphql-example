package graphql

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/jrkt/go-graphql-example/basic-plus/graphql/resolvers"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CompileSchema() (*graphql.Schema, error) {
	var s string

	err := filepath.Walk("./graphql/schema", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".graphql") {
			b, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			s += `
` + string(b)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return graphql.ParseSchema(s, resolvers.Init())
}
