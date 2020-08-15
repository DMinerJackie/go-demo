package main

import "sort"

func main() {

}

func intersect(nums1 []int, nums2 []int) []int {
	m0 := map[int]int{}
	for _, v := range nums1 {
		m0[v] += 1
	}

	k := 0
	for _, v := range nums2 {
		if m0[v] > 0 {
			m0[v] -= 1
			nums2[k] = v
			k++
		}
	}

	return nums2[0:k]
}

// 双指针解法
func intersect2(nums1 []int, nums2 []int) []int {
	i, j, k := 0, 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			nums1[k] = nums2[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}
