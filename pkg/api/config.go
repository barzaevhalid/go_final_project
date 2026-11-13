package api

import "os"

type Config struct {
	TODO_PASSWORD   string
	TODO_JWT_SECRET string
}

var Cfg Config

func Load() {
	Cfg = Config{
		TODO_PASSWORD:   os.Getenv("TODO_PASSWORD"),
		TODO_JWT_SECRET: os.Getenv("TODO_JWT_SECRET"),
	}
}
