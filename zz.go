package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
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

func main() {
	filename := "/Users/wang_qian0219/Desktop/group.jpg"
	EncodeImg(filename)
}
