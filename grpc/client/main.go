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
	//"io"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "../proto"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMicroClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 60)
	defer cancel()

	start := time.Now()
	// Obtener todos los gifs por stream
	/*stream, err := c.Top10Gifs(ctx, &pb.RequestFecha{Fecha: time.Now().Format("2006-01-02")})
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
		log.Printf("Server: %s", gif.Contenido)
	}*/
	// Obtener un solo gif por nombre
	gif, err := c.GetGif(ctx,  &pb.RequestGif{Fecha: time.Now().Format("2006-01-02"), Nombre: "giphy (8)"})
	end := time.Now()
	fmt.Println(end.Sub(start))
	log.Println(gif)
}

/*func Top10Gifs (c ) {
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
		log.Printf("Server: %s", gif.Contenido)
	}
}
*/