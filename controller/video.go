package controller

import (
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gocv.io/x/gocv"
)

var frame []byte
var mutex = &sync.Mutex{}

func Video(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=frame")
	data := ""
	for {
		mutex.Lock()

		data = "--frame\r\n  Content-Type: image/jpeg\r\n\r\n" + string(frame) + "\r\n\r\n"

		mutex.Unlock()

		time.Sleep(33 * time.Millisecond)

		c.Writer.Write([]byte(data))
	}
}
func Getframes() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatal("Error getting video device: ", err)
		return
	}
	defer webcam.Close()

	for {
		img := gocv.NewMat()
		// if img.Empty() {
		// 	continue
		// }
		webcam.Read(&img)

		buf, _ := gocv.IMEncode(".jpg", img)
		frame = buf.GetBytes()
	}
}
