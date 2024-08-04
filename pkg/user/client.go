package user

import (
	"log"

	pb "github.com/shivaraj-shanthaiah/godoc-API/pkg/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial() (pb.UserServiceClient, error) {
	gRPC, err := grpc.Dial(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing client on port: 8082")
		return nil, err
	}
	log.Printf("Successfully connected to client on 8082")
	return pb.NewUserServiceClient(gRPC), nil
}
