package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "cogno.com/grpc/alltypes"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	t := pb.NewAllTypesClient(conn)
	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	res, err := t.GetAllTypeData(ctx, &pb.RequestAllTypesData{
		Name:       "Shubham",
		Age:        24,
		Adult:      true,
		Info:       []byte("this is byte arr"),
		No1:        1,
		No2:        2,
		No3:        3,
		No4:        4,
		Fno1:       5,
		Fno2:       6,
		Fno3:       -10,
		Fno4:       -11,
		Day:        3,
		Weight:     80,
		Height:     5.9,
		HeightinCM: 175,
		Dict:       map[string]int32{"Shubham": 1, "Atish": 2},
		Array:      []int32{1, 2, 3, 4, 5, 6, 7},
	})
	if err != nil {
		log.Fatalf("could not get response: %v", err)
	}
	log.Printf("Response: %s", res.GetMsg())
}
