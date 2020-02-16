package main

import (
	"fmt"
	"learn-demo/day03_20200216_interface/real"
)

func main() {

	retriever := &real.Retriever{Contents: "1233"}

	contents := retriever.Get("https://y.qq.com/")

	//test timeout
	retriever.Get("https://www.google.com")

	fmt.Printf("the length is %d\n", len(contents))

}
