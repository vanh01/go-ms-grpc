package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	"go-grpc-ping/client"
	pong_pb "go-grpc-ping/client/pb"
	"go-grpc-ping/pb"
)

var pongClient pong_pb.PongClient

type pingServer struct {
	pb.UnimplementedPingServer
}

func (p pingServer) PingMessage(context context.Context, request *pb.PingRequest) (*pb.PingResponse, error) {
	fmt.Println(request.GetMessage())
	if len(request.GetMessage()) > 100 {
		return &pb.PingResponse{}, nil
	}
	message := "Ping response: " + request.GetMessage()
	pongClient.PongMessage(context, &pong_pb.PongRequest{Message: message})
	return &pb.PingResponse{}, nil
}

func main() {
	go func() {
		for pongClient == nil {
			time.Sleep(time.Second)
			pongClient = client.NewPongClient()
		}
		fmt.Println("succeed to connect to Pong Server")
	}()

	port := 9999
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pingServer := pingServer{}
	pb.RegisterPingServer(grpcServer, pingServer)
	grpcServer.Serve(lis)
}
