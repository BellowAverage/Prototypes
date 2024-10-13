package main

import (
	"errors"
	"fmt"
)

func main() {
	var first int32 = 1
	var second int32 = 2
	fmt.Println(swap(first, second))

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", result)
	}
}

func swap(first int32, second int32) (int32, int32) {
	return second, first
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
