package main

import (
	"log"
	"net"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "../proto"
	"google.golang.org/grpc/reflection"
	"github.com/go-redis/redis"
)

const (
	port = ":50051"
)

type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	val, err := client.HGetAll("gif").Result()
	if err != nil {
		fmt.Errorf("No se pudieron obrener los gifs &v", err)
	}
	fmt.Println(val)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
        return &pb.HelloReply{Message: "Hello again " + in.Name}, nil
}


func main() {


	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
