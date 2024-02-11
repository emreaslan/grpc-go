package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/emreaslan/grpc-go/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{4, 20, 100, 2, 8}

		for _, number := range numbers {
			log.Printf("Sendinh number: %d\n", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})

			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Problem while reading stream: %v\n", err)
				break
			}

			log.Printf("received a new maximum: %d\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
