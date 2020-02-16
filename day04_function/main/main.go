package main

import (
	"fmt"
	"learn-demo/day04_function/function"
)

func main() {

	// scene : 函数返回值可以为函数 : test adder function run
	adder := function.Adder()

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d\n", adder(i))
	}

	fmt.Printf("\n---------------{%s}----------------\n", "求和开始")

	// scene : 闭包 : test sum
	addSum := function.AdderSum(0)
	for i := 0; i <= 10; i++ {
		var sum int
		sum, addSum = addSum(i)
		fmt.Printf("%d\n", sum)
	}

	fmt.Printf("---------------{%s}----------------\n", "求和完毕")

	fmt.Printf("\n---------------{%s}----------------\n", "斐波那契")

	fibonacci := function.Fibonacci()

	for i := 0; i <= 10; i++ {
		fmt.Println(fibonacci())
	}

}
