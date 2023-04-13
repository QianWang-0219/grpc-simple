package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"grpc-simple/image"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	chunk_size = 1024
)

func EncodeImg(filename string) {
	//读原图片
	ff, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer ff.Close()
	sourcebuffer := make([]byte, 500000)
	n, _ := ff.Read(sourcebuffer)
	fmt.Println(n)
	//base64压缩
	sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])

	//写入临时文件
	//ioutil.WriteFile("a.png.txt", []byte(sourcestring), 0667)\
	ioutil.WriteFile("sum.txt", []byte(sourcestring), 0667)
}

func readIntFromCommendLine(reader *bufio.Reader, target *string) {
	_, err := fmt.Fscanf(reader, "%s\n", target)
	if err != nil {
		log.Fatalln("Cannot scan", err)
	}
}

func read_iterfile(path string, chunk_size int) {
	filename := filepath.Base(path)
	extension := filepath.Ext(path)
	fmt.Print("filename:", filename)
	fmt.Print("extension:", extension)

}

func main() {
	//filename := "/Users/wang_qian0219/Desktop/group.jpg"
	//EncodeImg(filename)
	// var path string
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("请输入待拼接图像文件夹地址: ")
	// readIntFromCommendLine(reader, &path)
	// //fmt.Println(path)

	// read_iterfile(path, chunk_size)
	//time.Sleep(60 * time.Second)
	fmt.Println(image.Image_mosaic("/Users/wang_qian0219/code/go/src/grpc-simple/resource"))

}
