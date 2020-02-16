package main

import (
	"fmt"
	"io/ioutil"
	"learn-demo/day01"
	"math"
	"math/cmplx"
	"strconv"
)

func main() {

	day01.PrintTime()

	var _ = 3
	var _ int = 3
	a := 4
	fmt.Println(a)

	exp := cmplx.Exp(1i*math.Pi) + 1
	fmt.Printf("%v\n", exp)

	c := 1i + 2
	fmt.Print(c)

	var sum int = switchType("+", 2, 3)
	fmt.Printf("%v\n", sum)

	// switchType("c",10,5)

	readFile("testFile")

	arrayTest()

	mapTest()

	node()

	number := 1234
	s := strconv.Itoa(number)
	fmt.Println(s)
	parseInt, _ := strconv.ParseInt(s, 10, 16)
	fmt.Println(parseInt)
}

func switchType(op string, a, b int) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	default:
		panic("illegal symbol" + op)

	}

	return result
}

func readFile(fileName string) {
	if content, err := ioutil.ReadFile(fileName); err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%s\n", content)
	}
}

func arrayTest() {
	// array is value type, not reference type; but in java language,it is reference type
	//sign variable
	var arry [5][4]int
	var arry3 [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(arry, arry3)
	// sing variable and sign value
	arry2 := [4]int{1, 2, 3, 4}
	fmt.Println(arry2)

	for i, v := range arry2 {
		fmt.Println(i, v)
	}

	for i := 0; i < len(arry2); i++ {
		fmt.Println(i, arry2[i])
	}
	//array type
	arry5 := [...]int{1334, 4444}
	fmt.Println(arry5)
}

func sliceTest() {
	arry2 := [4]int{1, 2, 3, 4}
	//slice
	_ = &arry2
	//slice
	_ = arry2[:]

	//slice: prn,len,cap

}

func mapTest() {
	map1 := map[string]string{
		"key1": "1",
		"ke2":  "2",
	}
	fmt.Println(map1)

	map2 := make(map[string]string) //empty map

	map2 = map[string]string{
		"key":  "1",
		"key2": "2",
	}

	var map3 map[string]string // empty map

	fmt.Println(map2, map3)

	if value, ok := map1["key1"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("the key is not exist")
	}

	delete(map1, "key2")

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	// slice , map ,function 的内建类型不能作为key，如果自定义struct不含前述内建类型，也可做key

}

func node() {

	//init
	var node day01.Node
	//populate
	node = day01.Node{Value: 1}

	node.Left = &day01.Node{Value: 2}

	node.Right = &day01.Node{Value: 3}

	//2
	node.Right.Right = new(day01.Node)

	fmt.Println(node.Value, node.Left.Value, node.Right.Value, node.Right.Right.Value)

	_ = []day01.Node{
		{1, nil, nil},
		{Value: 2},
		{3, nil, &node},
	}
}
