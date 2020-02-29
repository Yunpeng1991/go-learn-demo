package main

import (
	"bufio"
	"learn-demo/day06_20200218_errordeal/testserver"
	"os"
)

func main() {

	writeFile()
	testserver.TestServer()

}

func writeFile() {
	// 创建文件，写入文件，关闭文件
	fileName := "day06.txt"
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	//参数在defer语句时候计算，先记录后计算
	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush()

	writer.WriteString("i am a good boy")

}
