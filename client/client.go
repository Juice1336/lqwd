package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "test.com/lqwd_node/lqwd_node"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewNodeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var new_nodes = make(map[string]string)
	new_nodes["hk"] = "20.21.13.13"
	new_nodes["ca"] = "30.32.21.12"
	for name, ip := range new_nodes {
		r, err := c.SpawnNodes(ctx, &pb.NodeDetails{
			NodeName:  name,
			Ip:        ip,
			UserId:    "w3113123",
			Status:    pb.NodeDetails_RUNNING.Enum(),
			CreatedAt: "2020/11/11",
		})
		if err != nil {
			log.Fatalf("could not create nodes: %v", err)
		}
		log.Print(r)

	}
}
