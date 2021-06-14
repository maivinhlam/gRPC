package main

import (
	"calculator/calculatorpb"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Println("Sum called ...")
	resq := &calculatorpb.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}

	return resq, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("Err while create listen %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalcularotServiceServer(s, &server{})

	fmt.Println("Calculator is running ....")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("Err while serve &v", err)
	}
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PNDRequest, stream calculatorpb.CalcularotService_PrimeNumberDecompositionServer) error {
	log.Println("PrimeNumberDecomposition called ...")

	k := int32(2)
	N := req.GetNumber()
	for N > 1 {
		if N%k == 0 {
			N = N / k

			//send to client
			stream.Send(&calculatorpb.PNDResponse{
				Result: k,
			})

			time.Sleep(time.Second)
		} else {
			k++
			log.Printf("k increase to %v", k)
		}
	}
	return nil
}
