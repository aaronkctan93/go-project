package main

import "fmt"

func main() {

	// person := map[string]int{
	// 	"Mary": 180,
	// 	"John": 165,
	// 	"Emma": 170,
	// }
	// recordHigherIndex := 0
	temp := 0
	tempString := ""
	name := []string{"Mary", "John", "Emma"}
	height := []int{180, 165, 170}
	// nums := []int{2, 3, 4, 1, 6, 2}
	high := 0
	for i := 0; i < len(name)-1; i++ {
		high = height[i]
		for j := i + 1; j < len(name); j++ {
			if high < height[j] {
				high = height[j]
				temp = height[i]
				tempString = name[i]
				height[i] = height[j]
				name[i] = name[j]
				height[j] = temp
				name[j] = tempString
			}
		}
	}

	// for i := 0; i < len(nums)-1; i++ {
	// 	high = nums[i]
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if high < nums[j] {
	// 			high = nums[j]
	// 			temp = nums[j]
	// 			nums[j] = nums[i]
	// 			nums[i] = temp
	// 		}
	// 	}
	// 	fmt.Println(nums)
	// }

	fmt.Println(height)
	fmt.Println(name)
}

// https://leetcode.com/problems/sort-the-people/submissions/935142596/
