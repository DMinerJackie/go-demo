package main

import (
	"fmt"
	"github.com/pkg/errors"
)

var customedErrors = errors.New("invalid param error")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, customedErrors
	}

	return x / y, nil
}

func main() {
	fmt.Println(div(10, 5))
	fmt.Println(div(10, 0))
}
