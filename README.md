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