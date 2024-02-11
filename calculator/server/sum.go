package main

import (
	"context"
	"log"

	pb "github.com/emreaslan/grpc-go/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)

	return &pb.SumResponse{
		Result: in.First + in.Second,
	}, nil
}
