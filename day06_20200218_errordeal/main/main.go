package main

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"learn-demo/day06_20200218_errordeal/testserver"
	"os"
)

func main() {

	writeFile()

	testserver.TestServer()

	recoverDemo()

}

type userError interface {
	//identify message
	Message() string
	error
}

/**
  just like try-catch
*/
func recoverDemo() {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				//scene
				log.Warn("deal error", err.Error())
			} else {

				//scene 2
				// note : print recover , not err, because err is not recognized
				panic(fmt.Sprintf("i don`t konw this error %v", r))
			}
		}
	}()

	//tes scene 1
	//panic(errors.New("NullPointException"))

	//test scene 2
	panic(123)
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

//32 class
