package main

func main() {

}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
		} else {
			digits[i] = 0
		}
	}

	return append([]int{1}, digits...)
}
