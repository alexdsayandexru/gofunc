package gofunc

import "fmt"

func Version() {
	fmt.Println("3.0.0")
}

func Sum(ab ...int) int {
	b := 0
	for _, a := range ab {
		b = b + a
	}
	return b
}
