// main.go

package main

import (
  //"github.com/gin-gonic/contrib/static"
  "github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

  // Set the router as the default one provided by Gin
  router = gin.Default()
  gin.SetMode(gin.ReleaseMode)

  // Process the templates at the start so that they don't have to be loaded
  // from the disk again. This makes serving HTML pages very fast.
  router.LoadHTMLGlob("../public/templates/*")

    // Initialize the routes
  initializeRoutes()

	// Start serving the application
  router.Run(":3000")

}