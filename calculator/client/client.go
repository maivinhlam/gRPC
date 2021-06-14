package main

import (
	"context"
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

	callSum(client)
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
