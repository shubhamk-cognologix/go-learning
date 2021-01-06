// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "cogno.com/grpc/alltypes"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// ScalarTypesServer is used to implement all scalar types.
type ScalarTypesServer struct {
	pb.UnimplementedAllTypesServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
}

// GetAllTypeData : gets all types of data in request
func (s *ScalarTypesServer) GetAllTypeData(ctx context.Context, in *pb.RequestAllTypesData) (*pb.ResponseAllTypesData, error) {
	fmt.Println(in.GetName(), in.GetAdult(), in.GetInfo())
	fmt.Println(in.GetNo1(), in.GetNo2(), in.GetNo3(), in.GetNo4())
	fmt.Println(in.GetFno1(), in.GetFno2(), in.GetFno3(), in.GetFno4())
	fmt.Println(in.GetHeight(), in.GetHeightinCM(), in.GetWeight())
	fmt.Println(in.GetDay())
	fmt.Println(in.GetDict())
	fmt.Println(in.GetArray())
	return &pb.ResponseAllTypesData{Msg: "Success"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	pb.RegisterAllTypesServer(s, &ScalarTypesServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
