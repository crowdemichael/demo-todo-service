package services

import (
	"context"
	"time"

	"github.com/crowdeco/demo-todo-service/connections"
	"github.com/crowdeco/demo-todo-service/grpc/common"
	"github.com/crowdeco/demo-todo-service/grpc/user"
)

var UserService userServiceInterface = &userService{}

type userService struct {
	client user.UsersClient
}

func NewUserService() *userService {
	return &userService{
		client: connections.UserServiceConn(),
	}
}

type userServiceInterface interface {
	GetUserProfile(user.UsersClient, *user.IdInput) (*common.Response, error)
	CreateUser(user.UsersClient, *user.User) (*common.Response, error)
}

func (s *userService) GetUserProfile(client user.UsersClient, r *user.IdInput) (*common.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return client.GetUserProfile(ctx, r)
}

func (s *userService) CreateUser(client user.UsersClient, r *user.User) (*common.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return client.CreateUser(ctx, r)
}
