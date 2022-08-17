package img

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
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
	for {
		key := gocv.WaitKey(100)
		if 27 == key {
			fmt.Println("退出...")
			break
		}
	}

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

// ColorStyle 颜色转换
func ColorStyle(path string) {
	img := ReadImg(path)
	window := gocv.NewWindow("颜色转换")
	defer window.Close()
	window.IMShow(img)
	// 等待多少毫秒，为0则表示一直阻塞
	for {
		key := gocv.WaitKey(100)
		if 27 == key {
			fmt.Println("退出...")
			break
		}
		// 点击字母上的数字键 1的值为49
		if 49 == key {
			mat := gocv.NewMat()
			gocv.ApplyColorMap(img, &mat, gocv.ColormapAutumn)
			window.IMShow(mat)
		}
		// 点击字母上的数字键 1的值为49
		if 50 == key {
			mat := gocv.NewMat()
			// 颜色风格
			gocv.ApplyColorMap(img, &mat, gocv.ColormapBone)
			window.IMShow(mat)
		}
	}
}

// DrawRectangle 绘制矩形
func DrawRectangle() {
	// 创建空白的mat
	mat := gocv.Zeros(256, 256, gocv.MatTypeCV8UC3)
	// 创建矩形
	rect := image.Rect(50, 50, 100, 100)
	// 初始化颜色
	blue := color.RGBA{R: 1, G: 255}
	gocv.Rectangle(&mat, rect, blue, -1)
	ShowImg("矩形", mat)
}

// DrawCircle 绘制圆形
func DrawCircle() {
	// 创建空白的mat
	mat := gocv.Zeros(256, 256, gocv.MatTypeCV8UC3)
	defer mat.Close()
	// 右下角位置
	point := image.Point{X: 100, Y: 100}
	// 初始化颜色
	blue := color.RGBA{R: 1, G: 255}
	// 绘制
	gocv.Circle(&mat, point, 100, blue, -1)
	ShowImg("圆形", mat)
}

// DrawLine 绘制线条
func DrawLine() {
	// 创建空白的mat
	mat := gocv.Zeros(256, 256, gocv.MatTypeCV8UC3)
	defer mat.Close()
	// 左下角坐标
	p := image.Point{X: 10, Y: 10}
	// 右下角位置
	p1 := image.Point{X: 100, Y: 100}
	// 初始化颜色
	blue := color.RGBA{R: 1, G: 255}
	// 绘制
	gocv.Line(&mat, p, p1, blue, 1)
	ShowImg("直线", mat)
}

// DrawEllipse 绘制椭圆
func DrawEllipse() {
	// 创建空白的mat
	mat := gocv.Zeros(256, 256, gocv.MatTypeCV8UC3)
	defer mat.Close()
	// 长轴半径100，短轴半径50
	p := image.Point{X: 100, Y: 50}
	// 中心坐标
	center := image.Point{X: 100, Y: 100}
	// 初始化颜色
	blue := color.RGBA{R: 1, G: 255}
	// 绘制
	gocv.Ellipse(&mat, center, p, 90, 0, 360, blue, -1)
	ShowImg("椭圆", mat)
}

// DrawPolyLines 绘制多边形
func DrawPolyLines() {
	// 创建空白的mat
	mat := gocv.Zeros(256, 256, gocv.MatTypeCV8UC3)
	defer mat.Close()
	vector := gocv.NewPointsVector()
	pointVector := gocv.NewPointVector()
	pointVector.Append(image.Point{X: 100, Y: 50})
	pointVector.Append(image.Point{X: 100, Y: 100})
	pointVector.Append(image.Point{X: 50, Y: 100})
	vector.Append(pointVector)
	// 初始化颜色
	blue := color.RGBA{R: 1, G: 255}
	gocv.Polylines(&mat, vector, true, blue, 1)
	ShowImg("多边形", mat)
	mat2 := gocv.Zeros(256, 256, gocv.MatTypeCV8UC3)
	gocv.DrawContours(&mat2, vector, -1, blue, 1)
	ShowImg("多变形2", mat2)
}

// SplitAndMerge 通道分离及合并
func SplitAndMerge(path string) {
	img := ReadImg(path)
	defer img.Close()
	// 通道分离
	mats := gocv.Split(img)
	for i := range mats {
		mat := mats[i]
		var array = make([]gocv.Mat, 3)
		for j, _ := range array {
			if i == j {
				array[j] = mat
			} else {
				array[j] = gocv.Zeros(mat.Rows(), mat.Cols(), mat.Type())
			}
		}
		// 通道合并
		newMat := gocv.NewMat()
		gocv.Merge(array, &newMat)
		ShowImg("合并", newMat)
	}
}
