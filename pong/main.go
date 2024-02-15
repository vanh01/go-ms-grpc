package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	"go-grpc-pong/client"
	ping_pb "go-grpc-pong/client/pb"
	"go-grpc-pong/pb"
)

var pingClient ping_pb.PingClient = nil

type pongServer struct {
	pb.UnimplementedPongServer
}

func (p pongServer) PongMessage(context context.Context, request *pb.PongRequest) (*pb.PongResponse, error) {
	fmt.Println(request.GetMessage())
	message := "Pong response: " + request.GetMessage()
	pingClient.PingMessage(context, &ping_pb.PingRequest{Message: message})
	return &pb.PongResponse{}, nil
}

func main() {
	go func() {
		pingClient = client.NewPingClient()
		for pingClient == nil {
			time.Sleep(time.Second)
			pingClient = client.NewPingClient()
		}
		fmt.Println("succeed to connect to Ping Server")
		pingClient.PingMessage(context.Background(), &ping_pb.PingRequest{Message: "Test hehe"})
	}()

	port := 8888
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pongServer := pongServer{}
	pb.RegisterPongServer(grpcServer, pongServer)
	grpcServer.Serve(lis)

}
