package main

import (
	"context"
	"fmt"
	"log"
	"master_go_programming/gRPC/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()
	c := calculatorpb.NewSumServiceClient(cc)

	doUnary(c)
	doUnary(c)
	doUnary(c)
}

func doUnary(c calculatorpb.SumServiceClient) {
	fmt.Println("Starting unary rpc")
	req := &calculatorpb.SumRequest{
		A: 20,
		B: 30,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Reesponse from grreeet: %v", res.Sum)
}
