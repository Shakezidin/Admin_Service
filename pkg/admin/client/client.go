package client

import (
	"log"

	config "github.com/Shakezidin/config"
	pb "github.com/Shakezidin/pkg/admin/client/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.CoordinatorClient, error) {
	grpc, err := grpc.Dial(":"+cfg.COORDINATORAPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc client: %s, ", cfg.COORDINATORAPORT)
		return nil, err
	}
	log.Printf("succesfully Connected to coordinator Client at port: %v", cfg.COORDINATORAPORT)
	return pb.NewCoordinatorClient(grpc), nil
}
