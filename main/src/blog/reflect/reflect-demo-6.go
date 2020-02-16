package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := []int{1, 2, 3, 4}
	p1 := unsafe.Pointer(&a[1])
	p3 := unsafe.Pointer(uintptr(p1) + 2*unsafe.Sizeof(a[0]))
	*(*int)(p3) = 10
	fmt.Println(a)

	type Person struct {
		name   string
		age    int32
		gender bool
	}

	var i int8 = 1
	fmt.Println(i)
	var k uint8 = *(*uint8)(unsafe.Pointer(&i))
	fmt.Println(k)

	fmt.Println("reflect type: ", reflect.TypeOf(unsafe.Sizeof(a)))

	p := Person{name: "Jackie", age: 18, gender: true}
	fmt.Println("name alignof: ", unsafe.Alignof(p.name))     // type byte
	fmt.Println("age alignof: ", unsafe.Alignof(p.age))       // type int32
	fmt.Println("gender alignof: ", unsafe.Alignof(p.gender)) // type int64
	fmt.Println("person alignof: ", unsafe.Alignof(p))

	fmt.Println("name sizeof: ", unsafe.Sizeof(p.name))     // type byte
	fmt.Println("age sizeof: ", unsafe.Sizeof(p.age))       // type int32
	fmt.Println("gender sizeof: ", unsafe.Sizeof(p.gender)) // type int64
	fmt.Println("person sizeof: ", unsafe.Sizeof(p))
	fmt.Println("pointer person sizeof: ", unsafe.Sizeof(&p))

	pp := unsafe.Pointer(&p)
	pname := (*string)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(p.name)))
	page := (*int)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(p.age)))
	pgender := (*bool)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(p.gender)))
	*pname = "Stark"
	*page = 20
	*pgender = false

	fmt.Println(*pname, *page, *pgender)
}
