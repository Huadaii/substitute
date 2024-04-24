package src

import (
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"log"
	"os"

	"github.com/corona10/goimagehash"
)

// save *image.RGBA to filePath with PNG format.
func SaveImg(img *image.RGBA, filePath string) {
	if filePath == "" || img == nil {
		return
	}
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("SaveImg:", err)
	}
	defer file.Close()
	png.Encode(file, img)
}

// github.com/corona10/goimagehash
func Comparison(img1, img2 image.Image) int {
	if img1 == nil || img2 == nil {
		return 1000
	}
	hash1, err := goimagehash.AverageHash(img1)
	if err != nil {
		log.Println("image Comparison err:", err)
		return 1000
	}
	hash2, err := goimagehash.AverageHash(img2)
	if err != nil {
		log.Println("image Comparison err:", err)
		return 1000
	}
	distance, err := hash1.Distance(hash2)
	if err != nil {
		log.Println("image Comparison err:", err)
		return 1000
	}
	return distance
}

func ImgtoBinaryzation(img image.Image) image.Image {

	bounds := img.Bounds()
	binaryImg := image.NewGray(bounds) // 创建一个新的灰度图像
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			c := img.At(x, y)
			gray := color.GrayModel.Convert(c).(color.Gray)
			if gray.Y >= 128 {
				binaryImg.SetGray(x, y, color.Gray{uint8(255)}) // 大于等于128的像素点设为白色
			} else {
				binaryImg.SetGray(x, y, color.Gray{uint8(0)}) // 小于128的像素点设为黑色
			}
		}
	}

	return binaryImg
}
