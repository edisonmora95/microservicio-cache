package main

import (
	"log"
	"net"
	"fmt"
	"time"
	// "os"
	"encoding/json"
	//"strconv"

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

func (s *server) GetGif(ctx context.Context, in *pb.RequestFecha) (*pb.Gif, error) {
	start := time.Now()
	fmt.Println("Conectado al servidor GetGif")
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	val, _ := client.LRange(in.Fecha, 0, -1).Result()
	if len(val) == 0 {
		fmt.Println("No se encontró el gif en redis. Debe conectarse a mysql aquí.")
		end := time.Now()
		fmt.Println(end.Sub(start))
		return &pb.Gif{Titulo: "Primer gif"}, nil	
	}
	/*var gifs []Gif
	for index, element := range val {
		err := json.Unmarshal(element, &gifs)
	}*/
	fmt.Printf("Type of val: %T\n", val)
	end := time.Now()
	fmt.Println(end.Sub(start))
	/* if err != nil {
		fmt.Errorf("No se pudieron obrener los gifs &v", err)
		return nil, err
	} */
	/* contador, err := strconv.ParseInt(val["contador"], 10, 32)
	if err != nil {
		fmt.Errorf("No se pudo parsear el int &v", err)
		return nil, err
	} */
	return &pb.Gif{Titulo: "Hola"}, nil
}

func (s *server) Top10Gifs(in *pb.RequestFecha, stream pb.Micro_Top10GifsServer) error {
	// Primero se conecta a redis
	start := time.Now()
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	// Obtiene la lista de gifs
	val, err := client.LRange(in.Fecha, 0, -1).Result()
	if err != nil {
		fmt.Errorf("No se pudieron obrener los gifs &v", err)
		return err
	}
	if len(val) == 0 {
		fmt.Println("No se encontró el gif en redis. Debe conectarse a mysql aquí.")
		end := time.Now()
		fmt.Println(end.Sub(start))
		return nil	
	}
	var gif pb.Gif
	for _, gifStr := range val {
		err := json.Unmarshal([]byte(gifStr), &gif) // Los gifs se encuentran serializados en redis, por lo que hay que deserializar
		if err != nil {
			fmt.Println("There was an error:", err)
		}
		// Envia uno por uno los gifs deserializados por el stream
		if err := stream.Send(&gif); err != nil {
			return err
		}

	}
	return nil // Al finalizar le envia nil para indicarle al cliente que termino de enviar
}



func main() {
	fmt.Println("Servidor iniciado")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMicroServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
