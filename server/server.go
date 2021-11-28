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
func (s *server) GetNodesListByStatus(ctx context.Context, in *pb.Status) (*pb.Node, error) {
	n := &pb.Node{
		NodeName:  "aa",
		Ip:        "121.121.32.12",
		Status:    pb.Node_RUNNING.Enum(),
		CreatedAt: "2020/01/02",
		UserId:    "212121",
		Id:        "121121",
	}
	n1 := &pb.Node{
		NodeName:  "aWWWaa",
		Ip:        "121.121.32.12",
		Status:    pb.Node_FAILED.Enum(),
		CreatedAt: "2020/01/02",
		UserId:    "212SWS121",
		Id:        "sadadada",
	}
	if n1.Status == (*pb.Node_Status)(in.Status) {
		return n1, nil
	} else {
		return n, nil
	}
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
