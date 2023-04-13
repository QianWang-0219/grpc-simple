package encodeImage

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func EncodeImage(filename string) (string, error) {
	//读原图片
	//filename := "blockchain.PNG"
	ff, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer ff.Close()
	sourcebuffer := make([]byte, 500000)
	n, _ := ff.Read(sourcebuffer)
	//fmt.Println(n)
	//base64压缩
	sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])
	src := "data:image/PNG;base64," + sourcestring
	//fmt.Println(src)

	//写入临时文件
	//ioutil.WriteFile("a.png.txt", []byte(sourcestring), 0667)\
	ioutil.WriteFile("sum.txt", []byte(sourcestring), 0667)
	return src, nil
}

func DecodeImage() {
	//读取临时文件
	cc, _ := ioutil.ReadFile("a.png.txt")

	//解压
	dist, _ := base64.StdEncoding.DecodeString(string(cc))
	//写入新文件
	f, _ := os.OpenFile("xx.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(dist)
}
