package main

import (
	"bufio"
	"fmt"
	"learn-demo/day05_io/checkfile"
	"learn-demo/day05_io/ops"
	"strings"
	"time"
)

func main() {

	fmt.Printf("---------------{%s}---------------\n", "简单测试Scanner")
	reader := strings.NewReader("`123444444`" +
		"`122333333331`" +
		"`dddddddddd`")

	scanner := bufio.NewScanner(reader)

	if scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	fmt.Printf("---------------{%s}---------------\n", "简单测试Scanner完毕")

	file := &checkfile.CompareFile{}

	readFileData, _ := file.ReadFile("")

	readFileCompareData, _ := file.ReadFile("")

	file.Data = readFileData

	file.CompareData = readFileCompareData

	file.Compare()

	data := &ops.DiffData{
		Path:       "/yp/hello",
		DiffCount:  file.DiffCount,
		CreateTime: time.Now().UnixNano() / 1e6,
	}

	file.DifferData = data

	file.SaveDiff()

}
