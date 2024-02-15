package client

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go-grpc-pong/client/pb"
)

func NewPingClient() pb.PingClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:9999", opts...)
	if err != nil {
		log.Printf("failed to connect to the server: %v", err)
		return nil
	}
	if conn == nil {
		return nil
	}
	return pb.NewPingClient(conn)
}
