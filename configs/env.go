package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type env struct {
	AppPort  string
	GrpcPort string

	DbHost string
	DbPass string
	DbName string
	DbUser string
	DbPort string

	UserService string

	Debug bool
}

var (
	Env env
)

func init() {
	_ = godotenv.Load()

	Env.Debug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))

	Env.AppPort = os.Getenv("APP_PORT")
	Env.GrpcPort = os.Getenv("GRPC_PORT")

	Env.DbHost = os.Getenv("DB_HOST")
	Env.DbPass = os.Getenv("DB_PASS")
	Env.DbName = os.Getenv("DB_NAME")
	Env.DbUser = os.Getenv("DB_USER")
	Env.DbPort = os.Getenv("DB_PORT")

	Env.UserService = os.Getenv("USER_SERVICE")
}
