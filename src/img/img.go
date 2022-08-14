package img

import (
	"fmt"
	"gocv.io/x/gocv"
)

// ROI 绘制感兴趣的区域
func ROI() {

}

// ReadImg 读取图片
func ReadImg(path string) gocv.Mat {
	// 读取图片
	img := gocv.IMRead(path, gocv.IMReadUnchanged)
	if img.Empty() {
		panic("未成功加载图片...")
	}
	return img
}

// ShowImg 显示图片
func ShowImg(title string, img gocv.Mat) {
	window := gocv.NewWindow(title)
	defer window.Close()
	window.IMShow(img)
	// 等待多少毫秒，为0则表示一直阻塞
	gocv.WaitKey(0)
}

// ReadAndShowImg 读取并显示图片
func ReadAndShowImg(path string) {
	// 读取图片
	img := ReadImg(path)
	defer img.Close()
	ShowImg("显示", img)
}

// WriteImg 保存图像
func WriteImg(namePath string, img gocv.Mat) {
	// 第一个参数为图片许保存的路径及名称，如：./灰度图.jpeg
	ok := gocv.IMWrite(namePath, img)
	if ok {
		fmt.Println("图片保存成功")
	}
}

// ConvertColor 色彩转换
func ConvertColor(path string) {
	// 读取图片
	img := ReadImg(path)
	defer img.Close()

	// 定义灰度、hsv
	gray := gocv.NewMat()
	hsv := gocv.NewMat()
	// 转换为hsv
	gocv.CvtColor(img, &hsv, gocv.ColorBGRToHSV)
	// 显示图片
	ShowImg("hsv", hsv)
	// 转换为灰度图
	gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)
	// 显示图片
	ShowImg("gray", gray)
	// 保存灰度图片
	WriteImg("./灰度图.jpeg", gray)
}

// PixelVisit 像素处理
func PixelVisit(path string) {
	// 读取图片
	img := ReadImg(path)
	defer img.Close()
	//gray := gocv.NewMat()
	//defer gray.Close()
	//gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

	// 获取长、宽及通道数
	rows := img.Rows()
	cols := img.Cols()
	channels := img.Channels()

	// 创建空白的Mat对象
	newImg := gocv.Zeros(rows, cols, img.Type())
	defer newImg.Close()

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// 灰度只有1个通道
			if channels == 1 {
				// 获取并修改像素值
				pv := img.GetUCharAt(row, col)
				newImg.SetUCharAt(row, col, pv)
				continue
			}
			// 彩色为3通道
			if channels == 3 {
				// 获取并修改像素值
				vecb := img.GetVecbAt(row, col)
				for c := 0; c < channels; c++ {
					newImg.SetUCharAt(row, col*channels+c, vecb[c])
				}
			}
		}
	}

	// 增加亮度
	newImg.AddUChar(50)
	// 显示新图像
	ShowImg("新图像", newImg)
}
