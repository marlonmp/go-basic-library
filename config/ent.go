package config

import (
	_ "github.com/lib/pq"
	"github.com/marlonmp/go-basic-library/ent"
)

type EntConfig struct {
	Driver  string
	DSN     string
	Options []ent.Option
}

func Ent() EntConfig {

	return EntConfig{
		"postgres",
		"host=localhost port=5432 user=go-basic-library dbname=go-basic-library password=go-basic-library sslmode=disable",
		nil,
	}
}
