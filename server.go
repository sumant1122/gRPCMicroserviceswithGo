package main

import (
	"context"
	"log"
	"net"

	"github.com/micro/sumpb"
	"google.golang.org/grpc"
)

type server struct {
	sumpb.UnimplementedSumServer
}

func main() {
	endpoint := ":50051"
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	sumpb.RegisterSumServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server %v", err)
	}

}

// Add returns sum of two integers
func (*server) Add(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	a, b := req.GetNumbers().GetA(), req.GetNumbers().GetB()
	log.Println("A is : ", a)
	log.Println("B is : ", b)

	sum := a + b
	return &sumpb.SumResponse{Result: sum}, nil
}
