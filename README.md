# OpenCV学习

> OpenCV是一个基于Apache2.0许可（开源）发行的跨平台计算机视觉和机器学习软件库，可以运行在Linux、Windows、Android和Mac OS操作系统上。    
> 注：该学习内容使用gocv进行学习
## 图片处理API

### 读取图像`IMRead`及显示图像`IMShow`

> 注: `IMRead`方法返回的图像为 B,G,R

```go
package img

import (
	"fmt"
	"gocv.io/x/gocv"
)

const imgPath = "/home/huang/下载/09c60131bf1dee61dad8d6995d42059e.jpeg"

// ShowImg 读取并显示图片
func ShowImg() {
	// 读取图片
	img := gocv.IMRead(imgPath, gocv.IMReadUnchanged)
	if img.Empty() {
		fmt.Println("未成功加载图片...")
		return
	}
	window := gocv.NewWindow("显示")
	defer window.Close()
	defer img.Close()
	window.IMShow(img)
	// 等待多少毫秒，为0则表示一直阻塞
	gocv.WaitKey(0)
}
```

### 色彩转换`CvtColor`

```go
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
}
```

### 图像保存`IMWrite`
```go
// WriteImg 保存图像
func WriteImg(namePath string, img gocv.Mat) {
	// 第一个参数为图片许保存的路径及名称，如：./灰度图.jpeg
	ok := gocv.IMWrite(namePath, img)
	if ok {
		fmt.Println("图片保存成功")
	}
}
```

### 图像叠加`AddWeighted`

### 绘制矩形`Rectangle`
```go
// DrawRectangle 绘制矩形
func DrawRectangle() {
	// 创建空白的mat
	mat := gocv.Zeros(256, 256, gocv.MatTypeCV8UC3)
	// 创建矩形
	rect := image.Rect(50, 50, 100, 100)
	// 初始化颜色
	blue := color.RGBA{R: 1, G: 255}
	gocv.Rectangle(&mat, rect, blue, 1)
	ShowImg("矩形", mat)
}
```

### 绘制圆形`Circle`
```go
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
```

### 绘制线条`Line`
```go
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
```

### 绘制椭圆`Ellipse`
```go
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
```

### 绘制多边形`Polylines`或`DrawContours`
```go
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
```

### 通道分离`Split`和通道合并`Merge`
```go
// SplitAndMerge 通道分离及合并
func SplitAndMerge(path string) {
	img := ReadImg(path)
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
```

### 图像缩放`Resize`
```go
// ImgResize 图像缩放
func ImgResize(path string) {
	img := ReadImg(path)
	defer img.Close()
	// 缩小为一半
	size := image.Point{
		X: img.Cols() / 2,
		Y: img.Rows() / 2,
	}
	mat := gocv.NewMat()
	defer mat.Close()
	gocv.Resize(img, &mat, size, 0, 0, gocv.InterpolationLinear)
	fmt.Printf("%v, %v, %v, %v\n", img.Rows(), img.Cols(), mat.Rows(), mat.Cols())
	ShowImg("缩放", mat)
}
```

### 图像翻转`Flip`
```go
// ImgFlip 图像翻转
func ImgFlip(path string) {
    img := ReadImg(path)
    defer img.Close()
    mat := gocv.NewMat()
    defer mat.Close()
    // 0:上下翻转, 1:左右翻转，-1:对角线翻转
    gocv.Flip(img, &mat, 1)
    ShowImg("翻转", mat)
}
```

### 图像旋转`WarpAffine`
```go
// ImgRotate 图像旋转
func ImgRotate(path string) {
    img := ReadImg(path)
    defer img.Close()
    // 中心位置
    center := image.Point{
    X: img.Cols() / 2,
    Y: img.Rows() / 2,
    }
    mat := gocv.GetRotationMatrix2D(center, 45, 1)
    defer mat.Close()
    mat2 := gocv.NewMat()
    defer mat2.Close()
    size := image.Point{
    X: img.Cols(),
    Y: img.Rows(),
    }
    gocv.WarpAffine(img, &mat2, mat, size)
    ShowImg("旋转", mat2)
}
```