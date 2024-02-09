package userclient

import (
	"log"

	"github.com/Shakezidin/config"
	pb "github.com/Shakezidin/pkg/admin/userclient/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.UserClient, error) {
	grpc, err := grpc.Dial(":"+cfg.USERPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc client: %s, ", cfg.USERPORT)
		return nil, err
	}
	log.Printf("succesfully Connected to coordinator Client at port: %v", cfg.USERPORT)
	return pb.NewUserClient(grpc), nil
}
