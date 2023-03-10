package main

import (
	controller "RasperriPiV2Camera-Go/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	go controller.Getframes()

	server.Static("/assets", "./assets")
	server.Static("/node_modules", "./node_modules")

	server.LoadHTMLGlob("templates/index.html")

	server.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"status": http.StatusOK,
		})
	})

	server.GET("/video", controller.Video)

	server.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
