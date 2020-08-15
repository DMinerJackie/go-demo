package main

import "fmt"

func main() {
	binarySearch4()
}

//// 查找第一个值等于给定值的元素
//func binarySearch2() {
//	var arr = []int{1, 2, 3, 4, 4, 6}
//	hitValue := 4
//
//	low, high := 0, len(arr)-1
//	var mid int
//
//	for low <= high {
//		mid = low + (high-low)>>1
//		if arr[mid] >= hitValue {
//			high = mid - 1
//		} else {
//			low = mid + 1
//		}
//	}
//
//	if low < len(arr)-1 && arr[low] == hitValue {
//		fmt.Println("hit: ", low)
//		return
//	} else {
//		return
//	}
//}

// 查找第一个值等于给定值的元素
func binarySearch22() {
	var arr = []int{1, 2, 3, 4, 4, 4}
	hitValue := 4

	low, high := 0, len(arr)-1
	var mid int

	for low <= high {
		mid = low + (high-low)>>1
		if arr[mid] < hitValue {
			low = mid + 1
		} else if arr[mid] > hitValue {
			high = mid - 1
		} else {
			if mid == 0 || arr[mid-1] != hitValue { // 如果mid=0表示是第一个，那肯定是第一个；如果mid左边的mid-1不等于hitValue，也表示是第一个
				fmt.Println("hit: ", mid)
				return
			} else {
				high = mid - 1
			}
		}
	}
}

// 查找最后一个值等于给定值的元素
func binarySearch3() {
	var arr = []int{1, 2, 3, 4, 4, 4}
	hitValue := 4

	low, high := 0, len(arr)-1
	var mid int

	for low <= high {
		mid = low + (high-low)>>1
		if arr[mid] < hitValue {
			low = mid + 1
		} else if arr[mid] > hitValue {
			high = mid - 1
		} else {
			if mid == len(arr)-1 || arr[mid+1] != hitValue {
				fmt.Println("hit: ", mid)
				return
			} else {
				low = mid + 1
			}
		}
	}
}

// 查找第一个大于等于给定值的元素
func binarySearch4() {
	var arr = []int{1, 2, 3, 4, 5, 6}
	hitValue := 4

	low, high := 0, len(arr)-1
	var mid int

	for low <= high {
		mid = low + (high-low)>>1

		if arr[mid] >= hitValue {
			if mid == 0 || arr[mid-1] < hitValue {
				fmt.Println("hit: ", mid)
				return
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
}

// 查找第一个小于等于给定值的元素
func binarySearch5() {
	var arr = []int{1, 2, 3, 4, 5, 6}
	hitValue := 4

	low, high := 0, len(arr)-1
	var mid int

	for low <= high {
		mid = low + (high-low)>>1

		if arr[mid] <= hitValue {
			if mid == len(arr)-1 || arr[mid+1] > hitValue {
				fmt.Println("hit: ", mid)
				return
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}
	}
}

// 精确查找某个指定值
func binarySearch() {
	var arr = []int{1, 2, 3, 4, 5, 6}
	hitValue := 4

	low, high := 0, len(arr)-1
	var mid int

	for low <= high {
		mid = low + (high-low)>>1
		if arr[mid] == hitValue {
			fmt.Println("hit: ", mid)
			return
		} else if arr[mid] > hitValue {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
}
