package main

import "fmt"

type CustomError struct {
	x, y int
}

func (ce *CustomError) Error() string {
	return "invalid param error"
}

func division(x, y int) (int, error) {
	if y == 0 {
		return 0, &CustomError{x, y}
	}

	return x / y, nil
}

func main() {
	z, err := division(5, 0)
	if err != nil {
		switch e := err.(type) {
		case *CustomError:
			fmt.Println(e, e.x, e.y)
		default:
			fmt.Println(e)
		}
	}
	fmt.Println(z)
}
