package img

import (
	"math/rand"
	"testing"
)

const imgPath = "/home/huang/下载/09c60131bf1dee61dad8d6995d42059e.jpeg"

// 噪点图
const hotPixelImg = "/home/huang/图片/噪点图.jpeg"

func TestROI(t *testing.T) {
	ImageSmooth(hotPixelImg)
}

func TestCreateHotPixelImg(t *testing.T) {
	img := ReadImg(imgPath)
	defer img.Close()
	width := img.Cols()
	height := img.Rows()
	for i := 0; i < 5000; i++ {
		row := rand.Intn(height)
		col := rand.Intn(width)
		vecb := img.GetVecbAt(row, col)
		vecb[0] = 255
		vecb[1] = 255
		vecb[2] = 255
		for c := 0; c < 3; c++ {
			img.SetUCharAt(row, col*3+c, vecb[c])
		}
	}
	ShowImg("噪点图", img)
	WriteImg("/home/huang/图片/噪点图.jpeg", img)
}
