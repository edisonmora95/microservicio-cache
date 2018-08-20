package main

import (
	"log"
	"net"
	"fmt"
	"time"
	// "os"
	"encoding/json"
	//"strconv"
	"errors"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "../proto"
	"google.golang.org/grpc/reflection"
	"github.com/go-redis/redis"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

const (
	port = ":50051"
	//MYSQL constants
	mysqluser = "root"
	mysqlpassword = "root"
	mysqldatabase = "topgifs"
)

type server struct{}


func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
//returns top 10 gifs rows from MYSQL
func RetrieveGifsFromDB(user, password, database string) ([10]pb.Gif, error) {
	var gifs [10]pb.Gif
	db, err := sql.Open("mysql", user + ":" + password +"@/" + database)
	checkErr(err)
	// query: get me the top ten gifs on the database
    rows, err := db.Query("SELECT titulo, contenido, contador FROM gifs ORDER BY contador DESC LIMIT 10")
    checkErr(err)
    db.Close()
    i := 0
    for rows.Next() {
    	
            var titulo string
            var contenido string
            var contador int64
            err = rows.Scan(&titulo, &contenido, &contador)
            checkErr(err)
            gif :=  pb.Gif{}
            gif.Titulo = titulo
            gif.Contenido = contenido
            gif.Contador = contador
            fmt.Println(titulo)
            fmt.Println(contenido)
            fmt.Println(contador)
            gifs[i] = gif
            fmt.Println(gif)
            i++
        } 
    checkErr(err)

    return gifs, nil
}

//returns a gif row from MYSQL given the gif's title
func GetGifFromDB(user, password, database, titulo string) (pb.Gif, error){
	db, err := sql.Open("mysql", user + ":" + password +"@/" + database)
	checkErr(err)

	row, err := db.Query("SELECT titulo, contenido, contador FROM gifs where titulo = ?", titulo)
    checkErr(err)

    db.Close()
    var gif pb.Gif
    var title string
    var contenido string
    var contador int64
    err = row.Scan(&title, &contenido, &contador) 
    checkErr(err)
    gif.Titulo = titulo
    gif.Contenido = contenido
    gif.Contador = contador

    return gif, nil
}


func (s *server) GetGif(ctx context.Context, in *pb.RequestGif) (*pb.Gif, error) {
	start := time.Now()
	var gif pb.Gif
	var tempGif pb.Gif
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	val, err := client.LRange(in.Fecha, 0, -1).Result()
	if err != nil {
		fmt.Errorf("No se pudieron obrener los gifs &v", err)
		return &gif, err
	}
	if len(val) == 0 {
		fmt.Println("No se encontró el gif en redis. Debe conectarse a mysql aquí.")
		end := time.Now()
		fmt.Println(end.Sub(start))
		gif, err := GetGifFromDB(mysqluser, mysqlpassword, mysqldatabase,"hola")
		if err != nil {
			return nil, err
		}
		return &gif, nil	
	}	

	for _, gifStr := range val {
		err := json.Unmarshal([]byte(gifStr), &tempGif) // Los gifs se encuentran serializados en redis, por lo que hay que deserializar
		if err != nil {
			fmt.Println("There was an error:", err)
			return &gif, err
		}
		if gif.Titulo == in.Nombre {
			end := time.Now()
			fmt.Println(end.Sub(start))
			gif = tempGif
			return &gif, nil
		}
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
	return &gif, errors.new("Gif no encontrado")
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
		var gifs [10]pb.Gif
		fmt.Println("No se encontró el gif en redis. Debe conectarse a mysql aquí.")
		end := time.Now()
		fmt.Println(end.Sub(start))
		gifs, err = RetrieveGifsFromDB(mysqluser, mysqlpassword, mysqldatabase)
		for _, gif := range gifs {
			// Envia uno por uno los gifs por el stream
			if err := stream.Send(&gif); err != nil {
				return err
			}
		}
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
