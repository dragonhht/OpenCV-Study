package video

import (
	"fmt"
	"gocv.io/x/gocv"
)

// VideoCaptureDevice 调用摄像头
func VideoCaptureDevice() {
	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("预览")
	img := gocv.NewMat()

	defer webcam.Close()
	defer window.Close()
	defer img.Close()
	mat := gocv.NewMat()
	defer mat.Close()
	for {
		webcam.Read(&img)
		// 镜像处理
		gocv.Flip(img, &mat, 1)
		window.IMShow(mat)
		key := window.WaitKey(1)
		if key == 27 {
			break
		}
	}
}

// ReadVideo 读取视频文件
func ReadVideo(path string) {
	video, _ := gocv.VideoCaptureFile(path)
	mat := gocv.NewMat()
	window := gocv.NewWindow("播放")
	defer video.Close()
	defer mat.Close()
	defer window.Close()
	for {
		video.Read(&mat)
		window.IMShow(mat)
		key := window.WaitKey(30)
		if key == 27 {
			break
		}
	}
}

// VideoOperation 视频操作
func VideoOperation(path string) {
	video, _ := gocv.VideoCaptureFile(path)
	// 图像宽
	width := video.Get(gocv.VideoCaptureFrameWidth)
	// 图像高
	height := video.Get(gocv.VideoCaptureFrameHeight)
	// 总帧数
	count := video.Get(gocv.VideoCaptureFrameCount)
	// fps
	fps := video.Get(gocv.VideoCaptureFPS)
	// 编码格式
	fourcc := video.Get(gocv.VideoCaptureFOURCC)
	fmt.Printf("宽：%v, 高: %v, 总帧数: %v, fps: %v, 编码格式: %v\n", width, height, count, fps, fourcc)

	vw, _ := gocv.VideoWriterFile("./test.mp4", "MJPG", fps, int(width), int(height), true)
	defer vw.Close()

	mat := gocv.NewMat()
	window := gocv.NewWindow("播放")
	defer video.Close()
	defer mat.Close()
	defer window.Close()
	for {
		video.Read(&mat)
		window.IMShow(mat)
		// 写视频
		vw.Write(mat)
		key := window.WaitKey(24)
		if key == 27 {
			break
		}
	}
}
