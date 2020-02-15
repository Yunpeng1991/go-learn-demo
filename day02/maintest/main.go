package main

import (
	"fmt"
	"learn-demo/day02/retriever"
	"time"
)

func main() {
	fmt.Println("start------->", time.Now())
	sub := &retriever.Sub{123}
	resp := sub.Get("http://www.baidu.com")
	fmt.Println(resp)
	fmt.Println("end------->", time.Now())
}
