package image

import (
	"fmt"
	"io/ioutil"

	imgtype "github.com/shamsher31/goimgtype"
)

func GetFiles(folder string) []string {
	files, _ := ioutil.ReadDir(folder)

	var path string
	newSlice := make([]string, 0, 10)
	for _, file := range files {
		if file.IsDir() { // 用递归，获取目录以及子目录中所有文件
			GetFiles(folder + "/" + file.Name())
		} else {
			path = folder + "/" + file.Name()
			fmt.Println(folder + "/" + file.Name())
			// 获取图片的类型
			datatype, err2 := imgtype.Get(path)
			if err2 != nil {
				println(`不是图片文件`)
			} else {
				// 根据文件类型执行响应的操作
				switch datatype {
				case `image/jpeg`:
					println(`这是JPG文件`)
					newSlice = append(newSlice, path)
				case `image/png`:
					println(`这是PNG文件`)
				default:
					println(`这是其他文件`)
				}
			}
		}

	}
	return newSlice
}

// func main() {
// 	fmt.Println(GetFiles("./image"))
// }
