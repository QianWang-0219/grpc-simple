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

	"github.com/nfnt/resize"
)

const (
	MaxWidth = 3520
)

// 图片拼接（两张）
//
//	func MergeImageNew(base image.Image, mask image.Image) (*image.RGBA, error) {
//		width := base.Bounds().Max.X - 1 + mask.Bounds().Max.X - 1
//		height := base.Bounds().Max.Y - 1 + mask.Bounds().Max.Y - 1
//		dst := image.NewRGBA(image.Rect(0, 0, width, height))                                             // 创建一块画布
//		draw.Draw(dst, image.Rect(0, height/4, width/2, 3*height/4), base, image.Pt(0, 0), draw.Over)     // 绘制第一幅图
//		draw.Draw(dst, image.Rect(width/2, height/4, width, 3*height/4), mask, image.Pt(0, 0), draw.Over) // 绘制第二幅图
//		return dst, nil
//	}
func MergeImageNew(base image.Image, mask image.Image, paddingX int, paddingY int) (*image.RGBA, error) {
	base = resize.Resize(MaxWidth, 0, base, resize.Lanczos3)
	mask = resize.Resize(MaxWidth, 0, mask, resize.Lanczos3)
	baseSrcBounds := base.Bounds().Max

	maskSrcBounds := mask.Bounds().Max

	newWidth := baseSrcBounds.X
	newHeight := baseSrcBounds.Y

	maskWidth := maskSrcBounds.X
	maskHeight := maskSrcBounds.Y

	des := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight+maskHeight)) // 底板 newHeight+maskHeight 竖向排列
	//首先将一个图片信息存入jpg
	draw.Draw(des, des.Bounds(), base, base.Bounds().Min, draw.Over)
	//将另外一张图片信息存入jpg
	draw.Draw(des, image.Rect(paddingX, newHeight-paddingY+maskHeight, (paddingX+maskWidth), (newHeight-paddingY)), mask, image.ZP, draw.Over)

	return des, nil
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
	path := GetFiles(folder) // 返回所有jpg格式图片的路径
	var final_path string
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100)
	imgbase, err := GetImageFromFile(path[0])
	if err != nil {
		log.Fatalln(err)
	}
	for ix := 1; ix < len(path); ix++ {
		img, err := GetImageFromFile(path[ix])
		if err != nil {
			log.Fatalln(err)
		}
		imgbase, _ = MergeImageNew(imgbase, img, 0, 0)
	}
	final_path = "/Users/wang_qian0219/code/go/src/grpc-simple/mosaicRes/merge" + strconv.Itoa(num) + ".jpg"
	SaveImage(final_path, imgbase)
	return final_path
}
