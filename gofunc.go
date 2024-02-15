package gofunc

import "fmt"

func Version() {
	fmt.Println("1.0.0")
}

func Sum(a, b int) int {
	return a + b
}
