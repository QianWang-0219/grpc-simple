package utils

import (
	"path"
	"strings"
)

func ReadFilename(filePath string) (fileNameWithSuffix string, fileType string, fileNameOnly string) {
	//var filePath = "attachment/file/filename.txt"
	//获取文件名称带后缀
	fileNameWithSuffix = path.Base(filePath)
	//获取文件的后缀(文件类型)
	fileType = path.Ext(fileNameWithSuffix)
	//获取文件名称(不带后缀)
	fileNameOnly = strings.TrimSuffix(fileNameWithSuffix, fileType)
	// fmt.Printf("fileNameWithSuffix==%s\n fileType==%s;\n fileNameOnly==%s;",
	// 	fileNameWithSuffix, fileType, fileNameOnly)
	return fileNameWithSuffix, fileType, fileNameOnly
}
