package main

import (
		"fmt"
		"net/http"
		"os"
		"runtime"
		"path/filepath"
)

func retrieve_buff_gif(path string) []byte {

		// maximize CPU usage for maximum performance
		runtime.GOMAXPROCS(runtime.NumCPU())

		// open the uploaded file
		file, err := os.Open(path)

		if err != nil {
				fmt.Println(err)
				os.Exit(1)
		}

		buff := make([]byte, 512) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
		_, err = file.Read(buff)

		if err != nil {
				fmt.Println(err)
				os.Exit(1)
		}

		filetype := http.DetectContentType(buff)

		fmt.Println(filetype)

		defer file.Close()

		return buff


}

func getFilesScript(rootpath string) []string {

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
	if err != nil {
		fmt.Printf("walk error [%v]\n", err)
	}
	return list
}

func retrieve_all_gifs(folderPath string) [][]byte{
	listAll:= make([][]byte, 0,10)
	list := getFilesScript("gifs/")
	for _, p := range list {
		temp := retrieve_buff_gif(p)
		if temp != nil{
			listAll = append(listAll, temp)
		}
		
	}

	//fmt.Printf("size [%d]\n", len(list))

	return listAll
}

