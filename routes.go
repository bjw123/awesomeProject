package main

import (
	"awesomeProject/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		data := controller.MockGetAll()
		c.HTML(http.StatusOK, "hello-world.html", gin.H{
			"title": "Main website",
			"data":  data,
		})
	})

	router.GET("/post", func(c *gin.Context) {
		status := controller.MockPost()
		c.JSON(200, gin.H{
			"post request status": status,
		})
	})

	router.GET("/get", func(c *gin.Context) {
		//c.Request.URL.Query()
		data := controller.MockGetSelected(c.Query("id"))
		c.JSON(200, data)
	})

}
