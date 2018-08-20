// routes.go

package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
  )


func showIndexPage(c *gin.Context){
	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
		"title": "Home Page",
		},
	)	
}

func gifPage(c *gin.Context){
	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"gif.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
		"title": "Home Page",
		},
	)	
}


func tenGifsJson (c *gin.Context){

	//file := retrieve_all_gifs("gifs2/")
	file := retrieve_all_redis_gifs()
	// Call the HTML method of the Context to render a template
	c.JSON(
		http.StatusOK,
		gin.H{
		"counter": 3,
		"file": file,
		"timestamp": 35,
		},
	)	
}



func tenGifsMysqlJson (c *gin.Context){
	// Call the HTML method of the Context to render a template
	file := retrieve_all_gifs("gifs2/")
	// Call the HTML method of the Context to render a template
	c.JSON(
		http.StatusOK,
		gin.H{
		"counter": 3,
		"file": file,
		"timestamp": 35,
		},
	)		
}

func retrietOneGif (c *gin.Context){
	nameGif  := c.Param("name_gif")

	//functionToRetrietGif
	file := retrieve_one_gif(nameGif)

	c.JSON(
		http.StatusOK,
		gin.H{
		"counter": file.Contador,
		"file": file.Contenido,
		"name": file.Titulo,
		},
	)	



}
