package main

import (
	"fmt"
)

func main() {

	nums := []int{4, 7, 2, 3}
	largest := nums[0]
	recIndex := 0
	for i := 1; i < len(nums); i++ {
		if largest < nums[i] {
			largest = nums[i]
			recIndex = i
		}

		// fmt.Println(recIndex, largest)

	}
	// fmt.Println(recIndex)

	for i := recIndex; i < len(nums[recIndex+1:])+1; i++ {
		nums[i] = nums[i+1]
		fmt.Println(nums)
	}

	// fmt.Println(nums)
}
