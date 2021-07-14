package user

import (
	"context"
	"time"

	"github.com/crowdeco/demo-todo-service/grpc/common"
	"github.com/crowdeco/demo-todo-service/grpc/user"
)

var (
	UserService UserServiceInterface = &userService{}
)

type userService struct {
}

type UserServiceInterface interface {
	GetUserProfile(user.UsersClient, *user.IdInput) (*common.Response, error)
}

func (s *userService) GetUserProfile(client user.UsersClient, r *user.IdInput) (*common.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := client.GetUserProfile(ctx, r)
	if err != nil {
		return nil, err
	}

	return result, nil
}
