package main

import (
	"bufio"
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {

	writeFile()

	//testserver.TestServer()

	recoverDemo()

}

/**
  just like try-catch
*/
func recoverDemo() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			log.Warn("deal error", err.Error())
		}
	}()

	panic(errors.New("NullPointException"))
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
