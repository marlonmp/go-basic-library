package db

import (
	_ "github.com/lib/pq"

	"github.com/marlonmp/go-basic-library/config"
	"github.com/marlonmp/go-basic-library/ent"
)

var c *ent.Client

func Open(entConfig config.EntConfig) {

	client, err := ent.Open(entConfig.Driver, entConfig.DSN, entConfig.Options...)

	if err != nil {
		panic(err)
	}

	c = client
}

func GetClient() *ent.Client {
	return c
}
