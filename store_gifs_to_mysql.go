package main

import (
	"fmt"
	"runtime"
	"os"
	"strconv"
	"encoding/base64"
	"path/filepath"
)

import "database/sql"
import _ "github.com/go-sql-driver/mysql"



func RetrieveBuffGif(path string) string {
	// maximize CPU usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())
	// open the uploaded file
	file, err := os.Open(path)
	checkErr(err)

	buff := make([]byte, 1000000) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
	_, err = file.Read(buff)
	checkErr(err)

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
	checkErr(err)
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

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func main(){
	path := "./ginApp/server/gifsWithSize"
	list:= RetrieveAllGifs(path)

    db, err := sql.Open("mysql", "root:root@/topgifs")
	checkErr(err)

	for index, gif:=range list {
		fmt.Printf("ALMACENANDO EL GIF #%d\n", (index + 1))

        // insert
        stmt, err := db.Prepare("INSERT gifs SET titulo=?,contenido=?,contador=?")
        checkErr(err)


        res, err := stmt.Exec("gif-" + strconv.Itoa(index), gif, (index + 1))
        checkErr(err)

        //id, err := res.LastInsertId()
        //checkErr(err)

        fmt.Println(res)
        fmt.Println("Gif guardado en mysql")
        fmt.Println("***********************************************")
	}
        

}



