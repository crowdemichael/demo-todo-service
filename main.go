package main

import (
	"github.com/crowdeco/demo-todo-service/app"
	"github.com/crowdeco/demo-todo-service/database"
)

func init() {
	database.Connect()
}

func main() {
	app.Run()
}
