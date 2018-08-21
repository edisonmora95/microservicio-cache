/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

 package main

 import (
	 "log"
	 //"os"
	 "time"
	 "fmt"
	 "io"
 
	 "golang.org/x/net/context"
	 "google.golang.org/grpc"
	 pb "../../grpc/proto"
 )
 
 const (
	 address     = "localhost:50051"
	 defaultName = "world"
 )

 func retrieve_all_redis_gifs() []string {
	// Set up a connection to the server.
	listAll:= make([]string, 0,10)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMicroClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 60)
	defer cancel()

	start := time.Now()
	stream, err := c.Top10Gifs(ctx, &pb.RequestFecha{Fecha: time.Now().Format("2006-01-02")})
	end := time.Now()
	fmt.Println(end.Sub(start))
	if err != nil {
		log.Fatalf("could not get gif: %v", err)
	}
	for {
		gif, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
		}
		//log.Printf("Server: %s", gif)
		listAll = append(listAll, gif.Contenido)
	}

	return listAll
	

 }

 func retrieve_one_gif( str_id string) Gif{
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMicroClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 60)
	defer cancel()

	start := time.Now()
	//gif, err := c.GetGif(ctx,  &pb.RequestGif{Fecha: time.Now().Format("2006-01-02"), Nombre: "giphy (8)"})
	gif, err := c.GetGif(ctx,  &pb.RequestGif{Fecha: time.Now().Format("2006-01-02"), Nombre: str_id })
	end := time.Now()
	fmt.Println(end.Sub(start))
	return Gif{ Titulo: gif.Titulo, Contenido: gif.Contenido, Contador: gif.Contador}


 }


 
 
 