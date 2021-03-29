package app

import (
	"github.com/crowdeco/demo-todo-service/configs"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func Run() {
	mapUrls()
	port := configs.Env.AppPort
	router.Run(port)
}
