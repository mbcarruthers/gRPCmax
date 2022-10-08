package main

import (
	pb "github.com/mbcarruthers/gRPCmax/max/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	addr = "0.0.0.0:5051"
	tcp  = "tcp"
)

func main() {
	lis, err := net.Listen(tcp, addr)
	if err != nil {
		log.Fatalf("Error listening to %s:%s\n%s\n",
			tcp, addr, err.Error())
	}
	log.Println("Listening at 0.0.0.0:5051")
	s := grpc.NewServer()
	pb.RegisterMaxServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("grpc listen error at %s:%s \n %s\n",
			tcp, addr,
			err.Error())
	}

}
