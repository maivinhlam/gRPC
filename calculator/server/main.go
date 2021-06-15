package main

import (
	"calculator/proto"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *proto.SumRequest) (*proto.SumResponse, error) {
	log.Println("Sum called ...")
	resq := &proto.SumResponse{
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

	proto.RegisterCalculatorServiceServer(s, &server{})

	fmt.Println("Calculator is running ....")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("Err while serve &v", err)
	}
}

func (*server) PrimeNumberDecomposition(req *proto.PNDRequest, stream proto.CalculatorService_PrimeNumberDecompositionServer) error {
	log.Println("PrimeNumberDecomposition called ...")

	k := int32(2)
	N := req.GetNumber()
	for N > 1 {
		if N%k == 0 {
			N = N / k

			//send to client
			stream.Send(&proto.PNDResponse{
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

func (*server) Average(stream proto.CalculatorService_AverageServer) error {
	log.Println("Average called..")
	var total float32
	var count int
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			resp := &proto.AvgResponse{
				Result: total / float32(count),
			}

			return stream.SendAndClose(resp)
		}

		if err != nil {
			log.Fatalf("Err while Recv Average %v", err)
			return err
		}
		log.Printf("Receive num %v", req)
		total += req.GetNumber()
		count++
	}
}
