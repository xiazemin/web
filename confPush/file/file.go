package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type FileInfo struct {
	Name    string    `json:"name"`
	ModTime time.Time `json:"mod_time"`
}

func Exist(name string) bool {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		fmt.Println("file not exist!")
		return false
	}
	return true
}

func IsDir(name string) bool {
	fileInfo, _ := os.Stat(name)
	//是否是目录
	isDir := fileInfo.IsDir()
	fmt.Println(isDir)
	return isDir
}

func ModTime(name string) time.Time {
	fileInfo, _ := os.Stat(name)
	return fileInfo.ModTime()
}

func GetFiles(name string) []FileInfo {
	var names []FileInfo
	files, _ := ioutil.ReadDir(name)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			fmt.Println(file.Name())
			names = append(names, FileInfo{
				Name:    file.Name(),
				ModTime: file.ModTime(),
			})
		}
	}
	return names
}
