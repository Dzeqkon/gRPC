package server

import (
	pb "github.com/Dzeqkon/grpc-go-example/features/proto/echo"
	"golang.org/x/net/context"
	"log"
)

type Echo struct {
	pb.UnimplementedEchoServer
}

// UnaryEcho 一个普通的UnaryAPI
func (e *Echo) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("Recved: %v", req.GetMessage())
	resp := &pb.EchoResponse{Message: req.GetMessage()}
	return resp, nil
}
