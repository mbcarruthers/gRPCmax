package main

import (
	"context"
	"fmt"
	pb "github.com/mbcarruthers/gRPCmax/max/proto"
	"io"
	"log"
	"math/rand"
	"time"
)

func generateRandomRequests() []*pb.MaxRequest {
	var reqs []*pb.MaxRequest
	for i := 0; i < 20; i++ {
		value := int64(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100))
		reqs = append(reqs, &pb.MaxRequest{
			MaxReq: value,
		})
	}
	return reqs
}
func streamMax(c pb.MaxServiceClient) {
	log.Println("streamMax(client) called")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream client\n %s\n",
			err.Error())
	}
	requests := generateRandomRequests()
	wait := make(chan struct{})

	go func(reqSlice []*pb.MaxRequest) {
		for _, req := range reqSlice {
			if err = stream.Send(req); err != nil {
				log.Fatalf("Error sending request to server\n%s\n",
					err.Error())
			}
			time.Sleep(time.Millisecond * 750) //Todo: Remove this.
		}
		if err = stream.CloseSend(); err != nil {
			log.Fatalf("Erroring closing send stream\n%s\n",
				err.Error())
		}
	}(requests)

	go func() {
		for {
			if res, err := stream.Recv(); err == io.EOF {
				break
			} else if err != nil {
				log.Printf("Error receiving from stream\n%s\n",
					err.Error())
			} else {
				fmt.Printf("Result:%d\n", res.MaxRes)
			}
		}
		close(wait)
	}()
	<-wait
}
