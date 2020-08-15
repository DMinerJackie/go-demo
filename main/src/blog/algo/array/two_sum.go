package main

func main() {

}

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for k := i + 1; k < len(nums); k++ {
			if target-v == nums[k] {
				return []int{i, k}
			}
		}
	}
	return []int{}
}

func twoSum1(nums []int, target int) []int {
	result := []int{}

	m := make(map[int]int)
	for i, k := range nums {
		if value, exist := m[target-k]; exist {
			result = append(result, value)
			result = append(result, i)
		}

		m[k] = i
	}
	return result
}
