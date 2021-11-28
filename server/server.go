package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	pb "test.com/lqwd_node/lqwd_node"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedNodeServiceServer
}

func (s *server) SpawnNodes(ctx context.Context, in *pb.NodeDetails) (*pb.Node, error) {
	log.Printf("Received: %v", in.GetNodeName())
	var user_id string = string(rand.Intn(100))
	return &pb.Node{NodeName: in.GetNodeName(), Ip: in.GetIp(), Status: (*pb.Node_Status)(in.GetStatus().Enum()), UserId: in.GetUserId(), CreatedAt: in.GetCreatedAt(), Id: user_id}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNodeServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
