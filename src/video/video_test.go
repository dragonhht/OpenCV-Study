package video

import "testing"

const videoPath = "/home/huang/视频/190204084208765161.mp4"

// OpenCV人脸模型地址目录
const openCVFaceDNNRootPath = "/home/huang/apps/opencv-4.6.0/samples/dnn/face_detector"

func TestVideo(t *testing.T) {
	VideoOperation(videoPath)
}

// TestFaceDetection 人脸检测测试
func TestFaceDetection(t *testing.T) {
	FaceDetection("/home/huang/apps/opencv-4.6.0/data/haarcascades/haarcascade_frontalface_alt.xml")
}
