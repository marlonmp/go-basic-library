package config

import (
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type FiberConfig struct {
	Host   string
	Config fiber.Config
}

func Fiber() FiberConfig {

	config := fiber.Config{
		StrictRouting: true,
		CaseSensitive: true,
		AppName:       "go-basic-library",
		JSONEncoder:   jsoniter.Marshal,
		JSONDecoder:   jsoniter.Unmarshal,
	}

	return FiberConfig{
		"localhost:8000",
		config,
	}
}
