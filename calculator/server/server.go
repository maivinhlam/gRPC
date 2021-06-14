package main

import (
	"calculator/calculatorpb"
	"context"
	"fmt"
	"log"
	"net"

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
