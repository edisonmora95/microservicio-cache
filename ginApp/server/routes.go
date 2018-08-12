// routes.go

package main

func initializeRoutes() {

  // Handle the index route
  router.GET("/", showIndexPage)

  //Handle json response of 10 queries
  router.GET("/api/tenGifs", tenGifsJson)

  //Handle json response of 10 queries (mysql)
  router.GET("/api/tenGifsMysql", tenGifsMysqlJson)
  
}