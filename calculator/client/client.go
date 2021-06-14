package main

import (
	"context"
	"io"
	"log"

	"calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Err while dial %v", err)
	}

	defer cc.Close()

	client := calculatorpb.NewCalcularotServiceClient(cc)

	// log.Printf("Service client %f", client)

	// callSum(client)
	callPND(client)
}

func callSum(c calculatorpb.CalcularotServiceClient) {
	log.Println("Call sum api...")
	resp, err := c.Sum(context.Background(), &calculatorpb.SumRequest{
		Num1: 5,
		Num2: 7,
	})

	if err != nil {
		log.Fatalf("Call sum api err %v", err)
	}

	log.Printf("sum api response %v", resp.GetResult())
}

func callPND(c calculatorpb.CalcularotServiceClient) {
	log.Println("Call PND api...")

	stream, err := c.PrimeNumberDecomposition(context.Background(), &calculatorpb.PNDRequest{
		Number: 124,
	})

	if err != nil {
		log.Fatalf("Call PND error %v", err)
	}

	for {
		resp, recvErr := stream.Recv()
		if recvErr == io.EOF {
			log.Fatalf("Server finish streaming")
			return
		}
		log.Printf("PND api response: %v", resp.Result)
	}

}
