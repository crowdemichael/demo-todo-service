package connections

import (
	"log"

	"github.com/crowdeco/demo-todo-service/configs"
	"github.com/crowdeco/demo-todo-service/grpc/user"
	"google.golang.org/grpc"
)

func UserServiceConn() user.UsersClient {
	conn, err := grpc.Dial(configs.Env.UserService, grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	return user.NewUsersClient(conn)
}
