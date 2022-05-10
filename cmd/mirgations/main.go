//go:build ignore

package main

import (
	"context"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql/schema"

	"github.com/marlonmp/go-basic-library/config"
	"github.com/marlonmp/go-basic-library/pkg/db"
)

func main() {
	entConfig := config.Ent()

	db.Open(entConfig)

	client := db.GetClient()

	defer client.Close()

	// err := entc.Generate("./ent/schema", &gen.Config{}, entc.FeatureNames("sql/versioned-migration"))

	// if err != nil {
	// 	panic(err)
	// }

	ctx := context.Background()

	if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
		panic(err)
	}

	dir, err := migrate.NewLocalDir("./migrations")

	if err != nil {
		panic(err)
	}

	err = client.Schema.Diff(ctx, schema.WithDir(dir))

	if err != nil {
		panic(err)
	}
}
