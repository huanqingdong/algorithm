package swordo2

import "testing"

// 数组中的重复数字

func TestQ04(t *testing.T) {
	// println(" test q04")
	arr := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	target := 15

	x, y := len(arr), len(arr[0])
	for i, j := x-1, 0; i >= 0 && j < y; {

		v := arr[i][j]
		if v == target {
			print("true")
			return
		}

		if v < target {
			j++
		} else {
			i--
		}

	}

	print("false")

}
