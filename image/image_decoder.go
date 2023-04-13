package image

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// 图片拼接（两张）
func MergeImageNew(base image.Image, mask image.Image) (*image.RGBA, error) {
	width := base.Bounds().Max.X - 1 + mask.Bounds().Max.X - 1
	height := base.Bounds().Max.Y - 1 + mask.Bounds().Max.Y - 1
	dst := image.NewRGBA(image.Rect(0, 0, width, height))                                             // 创建一块画布
	draw.Draw(dst, image.Rect(0, height/4, width/2, 3*height/4), base, image.Pt(0, 0), draw.Over)     // 绘制第一幅图
	draw.Draw(dst, image.Rect(width/2, height/4, width, 3*height/4), mask, image.Pt(0, 0), draw.Over) // 绘制第二幅图
	return dst, nil
}

// 从本地读取图片
func GetImageFromFile(filePath string) (img image.Image, err error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err = jpeg.Decode(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read image from file:", filePath)
	fmt.Println(img.Bounds())
	return img, err
}

// 保存图片
func SaveImage(targetPath string, m image.Image) error {
	fSave, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer fSave.Close()

	err = jpeg.Encode(fSave, m, nil)

	if err != nil {
		return err
	}

	return nil
}

func Image_mosaic(folder string) string {
	path := GetFiles(folder)
	var final_path string
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100)

	imagA, err := GetImageFromFile(path[0])
	if err != nil {
		log.Fatalln(err)
	}
	imagB, err2 := GetImageFromFile(path[1])
	if err2 != nil {
		log.Fatalln(err)
	}
	img, err := MergeImageNew(imagA, imagB)
	if err != nil {
		log.Fatalln(err)
	}
	final_path = "/Users/wang_qian0219/code/go/src/grpc-simple/mosaicRes/merge" + strconv.Itoa(num) + ".jpg"
	SaveImage(final_path, img)
	return final_path
}
