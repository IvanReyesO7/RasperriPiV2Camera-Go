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
	server.LoadHTMLGlob("templates/index.html")

	server.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"status": http.StatusOK,
		})
	})

	server.GET("/video", controller.Video)

	server.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// func main() {
// 	webcam, err := gocv.VideoCaptureDevice(0)
// 	if err != nil {
// 		log.Fatal("Error getting video device: ", err)
// 		return
// 	}
// 	defer webcam.Close()
// 	stream := mjpeg.NewStream()

// 	go func() {
// 		for {
// 			img := gocv.NewMat()
// 			defer img.Close()
// 			webcam.Read(&img)

// 			buf, _ := gocv.IMEncode(".jpg", img)
// 			stream.UpdateJPEG(buf.GetBytes())
// 		}
// 	}()
// 	http.Handle("/", stream)
// 	err = http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		panic(err)
// 	}
// }
