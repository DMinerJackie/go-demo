package main

func myFunc() {
	i := 0
Here:
	println(i)
	i++
	if i < 1000 {
		goto Here
	}
}

func main() {
	myFunc()
}
