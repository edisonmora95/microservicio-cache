// routes.go

package main

func initializeRoutes() {

  // Handle the index route
  router.GET("/", showIndexPage)

  // Handle the gif route
  router.GET("/gif", gifPage)

  //Handle json response of 10 queries
  router.GET("/api/tenGifs", tenGifsJson)

  //Handle json response of 10 queries (mysql)
  router.GET("/api/tenGifsMysql", tenGifsMysqlJson)

  //Handle to receive one gif
  router.GET("/api/gif/:name_gif" , retrietOneGif)
  
}