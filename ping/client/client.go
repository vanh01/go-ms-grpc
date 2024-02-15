package client

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go-grpc-ping/client/pb"
)

func NewPongClient() pb.PongClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:8888", opts...)
	if err != nil {
		log.Printf("failed to connect to the server: %v", err)
		return nil
	}
	if conn == nil {
		return nil
	}
	return pb.NewPongClient(conn)
}
