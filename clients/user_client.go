package clients

import (
	"log"
	"os"

	"github.com/crowdeco/demo-todo-service/grpc/user"
	"google.golang.org/grpc"
)

func UserConn() user.UsersClient {
	port := os.Getenv("USER_SERVICE")
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}
	return user.NewUsersClient(conn)
}
