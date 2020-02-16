package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

	reader := strings.NewReader("123444444")

	scanner := bufio.NewScanner(reader)

	if scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
