package main

import (
	pb "github.com/mbcarruthers/gRPCmax/max/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	addr = "localhost:5051"
)

func main() {
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error dialing %s:\n%s\n",
			addr, err.Error())
	}
	defer conn.Close()

	maxClient := pb.NewMaxServiceClient(conn)

	streamMax(maxClient)
}
