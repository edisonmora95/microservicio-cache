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

const (
	//MYSQL constants modify as needed
	mysqluser = "root"
	mysqlpassword = "root"
	mysqldatabase = "topgifs"
)


func createDB(){

   db, err := sql.Open("mysql", mysqluser + ":" + mysqlpassword +"@tcp(127.0.0.1:3306)/")
   if err != nil {
       panic(err)
   }
   defer db.Close()

   _,err = db.Exec("CREATE DATABASE " + mysqldatabase)
   if err != nil {
       panic(err)
   }

   _,err = db.Exec("USE " + mysqldatabase)
   if err != nil {
       panic(err)
   }

   _,err = db.Exec("Create table gifs (id int(10) key not null auto_increment, titulo VARCHAR(64) not null, contenido longblob not null, contador BIGINT not null )")
   if err != nil {
       panic(err)
   }
}

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

	createDB()
	path := "./ginApp/server/gifsWithSize"

	list:= RetrieveAllGifs(path)

    db, err := sql.Open("mysql", mysqluser + ":" + mysqlpassword +"@/" + mysqldatabase)
	checkErr(err)

	for index, gif:=range list {
		fmt.Printf("ALMACENANDO EL GIF #%d\n", (index + 1))

        // insert
        stmt, err := db.Prepare("INSERT gifs SET titulo=?,contenido=?,contador=?")
        checkErr(err)


        res, err := stmt.Exec("gif-" + strconv.Itoa(index), gif, (index + 1))
        checkErr(err)

        id, err := res.LastInsertId()
        checkErr(err)

         fmt.Println(id)
        fmt.Println("Gif guardado en mysql")
        fmt.Println("***********************************************")
	}
        

}



