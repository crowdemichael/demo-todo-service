package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/crowdeco/demo-todo-service/clients"
	grpc "github.com/crowdeco/demo-todo-service/grpc/user"
	"github.com/crowdeco/demo-todo-service/services/user"
	"github.com/gin-gonic/gin"
)

type UserProfile struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUserProfile(c *gin.Context) {
	id := c.Param("id")
	userId, _ := strconv.ParseInt(id, 10, 64)

	gres, err := user.UserService.GetUserProfile(clients.UserConn(), &grpc.IdInput{Id: userId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	res := UserProfile{}
	json.Unmarshal(gres.Data, &res)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   res,
	})
}
