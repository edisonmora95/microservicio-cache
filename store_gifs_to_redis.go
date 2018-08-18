package main

import (
	"fmt"
	"runtime"
	"time"
	"os"
	"strconv"
	"encoding/json"
	"encoding/base64"
	"path/filepath"

	"github.com/go-redis/redis"
)

// Esta funcion maneja los errores
func check(e error) {
	if e != nil {
		panic(e)
		os.Exit(1)
	}
}

func RetrieveBuffGif(path string) string {
	// maximize CPU usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())
	// open the uploaded file
	file, err := os.Open(path)
	check(err)

	buff := make([]byte, 70000) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
	_, err = file.Read(buff)
	check(err)

	defer file.Close()

	imgBase64Str := base64.StdEncoding.EncodeToString(buff)

	return imgBase64Str
}

func GetFilesScript(rootpath string) []string {
	list := make([]string, 0, 10)

	err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".gif" {
			list = append(list, path)
		}
		return nil
	})
	check(err)
	return list
}

func RetrieveAllGifs(folderPath string) []string{
	listAll:= make([]string, 0,10)
	list := GetFilesScript(folderPath)
	for _, p := range list {
		temp := RetrieveBuffGif(p)
		listAll = append(listAll, temp)
	}
	return listAll
}

type Gif struct {
	Titulo string
	Contenido string
	Contador int64
}

func main() {
	today := time.Now().Format("2006-01-02")
	path := "./gifs"
	list := RetrieveAllGifs(path)
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	fmt.Println("Cliente conectado")
	for index, gif := range list {
        fmt.Printf("ALMACENANDO EL GIF #%d\n", (index + 1))
        nombre := "gif-" + strconv.Itoa(index)
        gifStruct := &Gif{Titulo: nombre, Contenido: gif, Contador: int64(index+1)}
        gifMarshal, err := json.Marshal(gifStruct)
        check(err)
        client.LPush(today, gifMarshal)
        /*client.HMSet(nombre, map[string]interface{}{
        	"contenido": gif,
        	"contador": float64(index+1),
    	})*/
        /* client.ZAdd("gifs", redis.Z{
			Score: float64(index+1),
			Member: gif,
		}) */
		fmt.Println("Gif guardado en redis")
        fmt.Println("***********************************************")
    }
}