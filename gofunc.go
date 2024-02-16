package gofunc

import "fmt"

func Version() {
	fmt.Println("2.0.0")
}

func Sum(abc ...int) int {
	b := 0
	for _, a := range abc {
		b = b + a
	}
	return b
}
