package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	getRouter()
}

func getRouter() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*")
	setRoutes(router)
	log.Println("Open http://localhost:8080 in the browser")
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
