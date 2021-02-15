package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"master_go_programming/gRPC/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)

	doServerStreaming(c)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting streaming server rpc")
	req := &greetpb.GreetManyTimesRequest{Greeting: &greetpb.Greeting{
		FirstName: "Igna",
		LastName:  "Darius",
	}}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("rresponse from greeat many times :%v", msg.GetResult())
	}

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting unary rpc")
	req := &greetpb.GreatRequest{Greeting: &greetpb.Greeting{
		FirstName: "igna",
		LastName:  "darius",
	}}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Reesponse from grreeet: %v", res.Result)
}
