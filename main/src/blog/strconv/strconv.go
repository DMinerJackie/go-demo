package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("============AppendBool============")
	b := []byte("bool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b))

	fmt.Println("============AppendFloat============")
	b32 := []byte("float32:")
	b32 = strconv.AppendFloat(b32, 3.1415926535, 'E', -1, 32) // -1所在标识表示精度，-1表示尽可能显示更多来确保准确
	fmt.Println(string(b32))

	b64 := []byte("float64:")
	b64 = strconv.AppendFloat(b64, 3.1415926535, 'E', -1, 64)
	fmt.Println(string(b64))

	fmt.Println("============AppendInt============")
	b10 := []byte("int (base 10):")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))

	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))

	fmt.Println("============AppendQuote============")
	b1 := []byte("quote:")
	b1 = strconv.AppendQuote(b1, `"Fran & Freddie's Diner"`)
	fmt.Println(string(b1))

	fmt.Println("============AppendQuoteRune============")
	b2 := []byte("rune:")
	b2 = strconv.AppendQuoteRune(b2, '☺')
	fmt.Println(string(b2))

	fmt.Println("============AppendQuoteRuneToASCII============")
	b3 := []byte("rune (ascii):")
	b3 = strconv.AppendQuoteRuneToASCII(b3, '☺')
	fmt.Println(string(b3))

	fmt.Println("============AppendUint============")
	b4 := []byte("uint (base 10):")
	b4 = strconv.AppendUint(b4, 42, 10)
	fmt.Println(string(b10))

	b5 := []byte("uint (base 16):")
	b5 = strconv.AppendUint(b5, 42, 16)
	fmt.Println(string(b16))

	fmt.Println("============Atoi============")
	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	/**
	format表示格式化，将所有类型都格式化为string类型
	object(int, bool, float) -> string
	parse表示解析，即从字符串类型解析成对应的类型
	string -> object
	*/
	fmt.Println("============FormatBool============")
	v2 := true
	s2 := strconv.FormatBool(v2)
	fmt.Printf("%T, %v\n", s2, s2)

	fmt.Println("============FormatFloat============")
	v3 := 3.1415926535

	s32 := strconv.FormatFloat(v3, 'E', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32)

	s64 := strconv.FormatFloat(v3, 'E', -1, 64)
	fmt.Printf("%T, %v\n", s64, s64)

	fmt.Println("============FormatInt============")
	v4 := int64(-42)

	s10 := strconv.FormatInt(v4, 10)
	fmt.Printf("%T, %v\n", s10, s10)

	s16 := strconv.FormatInt(v4, 16)
	fmt.Printf("%T, %v\n", s16, s16)

	fmt.Println("============Itoa============")
	i := 10
	s := strconv.Itoa(i)
	fmt.Printf("%T, %v\n", s, s)

	fmt.Println("============FormatInt============")
	v5 := int64(-42)

	s101 := strconv.FormatInt(v5, 10)
	fmt.Printf("%T, %v\n", s101, s101)

	s161 := strconv.FormatInt(v5, 16)
	fmt.Printf("%T, %v\n", s161, s161)

	fmt.Printf("%f\n", math.Pow10(8))

}
