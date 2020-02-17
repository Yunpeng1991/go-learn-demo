package main

import (
	"bufio"
	"fmt"
	"strings"
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

}
