package main

import "fmt"

func main() {
	i := 10

	/**
	Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，
	而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码。
	*/
	switch i {
	case 1:
		fmt.Println("1=1")
	case 10:
		fmt.Println("=10")
		fallthrough
	case 2, 3, 4:
		fmt.Println("=2,3,4")
		fallthrough
	default:
		fmt.Println("unknown")
	}

	fmt.Println("==========================")

	integer := 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case 8:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

}
