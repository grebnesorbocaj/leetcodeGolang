package main

import "fmt"

func main() {
	input := []int{2, 7, 15, 21}
	target := 9
	res := twoSum(input, target)

	if res[0] == 0 && res[1] == 1 {
		fmt.Println("success")
	}
}

func twoSum(nums []int, target int) []int {
	var x, y int
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]+nums[i] == target {
				x = i
				y = j
				break
			}
		}
	}
	return []int{x, y}
}
