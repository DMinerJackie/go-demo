package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	vv, ok := m.LoadOrStore("1", "one")

	fmt.Println(vv, ok) //one false

	vv, ok = m.Load("1")

	fmt.Println(vv, ok) //one true

	vv, ok = m.LoadOrStore("1", "oneone")

	fmt.Println(vv, ok) //one true

	m.Store("1", "oneone")

	vv, ok = m.Load("1")

	fmt.Println(vv, ok) // oneone true

	m.Store("2", "two")

	m.Range(func(k, v interface{}) bool {

		fmt.Println(k, v)

		return true

	})
}
