package video

import "gocv.io/x/gocv"

// VideoCaptureDevice 调用摄像头
func VideoCaptureDevice() {
	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
