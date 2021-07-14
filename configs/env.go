package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	AppPort string

	UserServiceGrpc string
}

var (
	Env env
)

func init() {
	_ = godotenv.Load()

	Env.AppPort = os.Getenv("APP_PORT")

	Env.UserServiceGrpc = os.Getenv("USER_SERVICE")
}
