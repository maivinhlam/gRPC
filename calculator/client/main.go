package main

import (
	"context"
	"io"
	"log"
	"time"

	"calculator/proto"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Err while dial %v", err)
	}

	defer cc.Close()

	client := proto.NewCalculatorServiceClient(cc)

	// log.Printf("Service client %f", client)

	// callSum(client)
	// callPND(client)
	callAverage(client)
}

func callSum(c proto.CalculatorServiceClient) {
	log.Println("Call sum api...")
	resp, err := c.Sum(context.Background(), &proto.SumRequest{
		Num1: 5,
		Num2: 7,
	})

	if err != nil {
		log.Fatalf("Call sum api err %v", err)
	}

	log.Printf("sum api response %v", resp.GetResult())
}

func callPND(c proto.CalculatorServiceClient) {
	log.Println("Call PND api..")

	stream, err := c.PrimeNumberDecomposition(context.Background(), &proto.PNDRequest{
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

func callAverage(c proto.CalculatorServiceClient) {
	log.Println("Call Average api..")
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Call average err %v", err)
	}

	listReq := []proto.AvgRequest{
		proto.AvgRequest{
			Number: 4,
		},
		proto.AvgRequest{
			Number: 6,
		},
		proto.AvgRequest{
			Number: 8,
		},
		proto.AvgRequest{
			Number: 12,
		},
		proto.AvgRequest{
			Number: 10,
		},
		proto.AvgRequest{
			Number: 14,
		},
	}

	for _, req := range listReq {
		err := stream.Send(&req)
		if err != nil {
			log.Fatalf("send average request err %v", err)
		}
		time.Sleep(time.Second)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("receive average response err %v", err)
	}

	log.Fatalf("average response %+v", resp)

}
