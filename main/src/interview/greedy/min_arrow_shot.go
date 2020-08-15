package main

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	end := points[0][1]
	var result int
	for i := 0; i < len(points); i++ {
		if points[i][0] < end {
			result++
		} else {
			end = points[i][1]
		}
	}
	return result
}
