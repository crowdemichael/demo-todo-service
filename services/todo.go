package services

import (
	"encoding/json"

	"github.com/crowdeco/demo-todo-service/connections"
	"github.com/crowdeco/demo-todo-service/domain/user"
	grpc "github.com/crowdeco/demo-todo-service/grpc/user"
	"github.com/jinzhu/copier"
)

func GetUserProfile(id int64) (*user.UserProfile, error) {
	payload, err := UserService.GetUserProfile(connections.UserServiceConn(), &grpc.IdInput{Id: id})
	if err != nil {
		return nil, err
	}

	res := user.UserProfile{}
	json.Unmarshal(payload.Data, &res)

	return &res, nil
}

func CreateUser(value *user.UserProfile) (*user.UserProfile, error) {
	req := grpc.User{}
	copier.Copy(&req, &value)

	payload, err := UserService.CreateUser(connections.UserServiceConn(), &req)
	if err != nil {
		return nil, err
	}

	res := user.UserProfile{}
	json.Unmarshal(payload.Data, &res)

	return &res, nil
}
