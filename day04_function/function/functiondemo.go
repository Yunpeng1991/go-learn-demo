package function

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//very important
/*return a adder function (闭包)**/
func Adder() func(int) int {
	// sum is a local variable(不能有状态: very important
	sum := 0
	return func(i int) int {
		//sum = i+sum  sum variable will be summation: very important
		// i (自由变量)
		return sum + i
	}
}

//:定义一个func类型
type iAdder func(int) (int, iAdder)

/*sum 递归 累加**/
func AdderSum(base int) func(int) (int, iAdder) {
	//v 为自由变量，即main方法里传入的变量
	return func(v int) (int, iAdder) {
		return base + v, AdderSum(base + v)
	}
}

//斐波那契
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//

func Fibonacci2() GenFibonacci {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type GenFibonacci func() int

func (gen GenFibonacci) Read(data []byte) (i int, err error) {
	next := gen()
	if next > 10000 {
		//stop scan
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	//Read(data) 将data写入(gen GenFibonacci) Read(data []byte)(i int,err error)
	// if p is small ,will has a bug, nee a buffer to resolve
	return strings.NewReader(s).Read(data)
}

func PrintFileContents(read io.Reader) {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
