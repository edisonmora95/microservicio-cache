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