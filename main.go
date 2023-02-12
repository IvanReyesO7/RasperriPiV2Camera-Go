package main

import (
	"log"
	"net/http"

	"github.com/hybridgroup/mjpeg"
	"gocv.io/x/gocv"
)

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatal("Error getting video device: ", err)
		return
	}
	defer webcam.Close()
	stream := mjpeg.NewStream()

	go func() {
		for {
			img := gocv.NewMat()
			defer img.Close()
			webcam.Read(&img)

			buf, _ := gocv.IMEncode(".jpg", img)
			stream.UpdateJPEG(buf.GetBytes())
		}
	}()
	http.Handle("/", stream)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
