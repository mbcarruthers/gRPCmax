package main

import (
	"fmt"
	pb "github.com/mbcarruthers/gRPCmax/max/proto"
	"io"
	"log"
	"sort"
)

type Server struct {
	pb.MaxServiceServer
}

func biggest(set []int64) int64 {
	sort.Slice(set, func(i, j int) bool {
		return set[i] > set[j]
	})
	return set[0]
}

func (s *Server) Max(stream pb.MaxService_MaxServer) error {
	log.Printf("Server side max invoked\n")
	var valueSet []int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error receiving stream:\n%s\n",
				err.Error())
		} else {
			value := req.MaxReq
			fmt.Printf("Value sent from client %d\n", value)
			valueSet = append(valueSet, value)
			max := biggest(valueSet)
			if err = stream.Send(&pb.MaxResponse{
				MaxRes: max,
			}); err != nil {
				log.Fatalf("Error streaming data to client\n%s\n",
					err.Error())
			}
		}
	}
	return nil
}
