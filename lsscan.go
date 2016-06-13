package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//파라메터를 안 받았다면 현재 경로 "." 이하 파일, 디렉토리들 출력
	if len(os.Args) == 1 {
		listDir(".")
	} else {
		listDir(os.Args[1])
	}
}

func listDir(dirFile string) {
	files, _ := ioutil.ReadDir(dirFile)
	for _, f := range files {
		t := f.ModTime()
		fmt.Println(f.Name(), dirFile+"/"+f.Name(), f.IsDir(), t, f.Size()) //파일(디렉토리)명, 전체path, 디렉토리여부, 생성시간, 파일크기 출력
		if f.IsDir() {
			listDir(dirFile + "/" + f.Name())
		}
	}
}
