package main

import "fmt"

func main() {

	// nums := []int{5, 7, 1, 10, 2}
	students := []int{2, 7, 4}
	seats := []int{3, 1, 5}
	countMove := 0
	checkIsSmallest := 0
	recordPosition := 0
	distance := 0

	for _, student := range students {
		for k, seat := range seats {
			distance = student - seat
			if distance < 0 {
				distance *= -1
			}
			if k == 1 {
				checkIsSmallest = distance
			} else {
				if seat != -1 && checkIsSmallest > distance {
					checkIsSmallest = distance
					recordPosition = k
				}
			}
		}
		fmt.Println(seats)
		if seats[recordPosition] != -1 {
			seats[recordPosition] = -1
			// fmt.Println(recordPosition)
		}

		countMove += checkIsSmallest
	}
	fmt.Println(countMove)
	// fmt.Println(findBiggest(students, seats))
}

/*
// 找最大
func findBiggest(students, seats []int) int {
	appendFindBiggest := appendSlice(students, seats)

	bigNum := appendFindBiggest[0]

	for i, num := range appendFindBiggest {
		if i != 0 && num > bigNum {
			bigNum = num
		}
	}

	return bigNum
}

func appendSlice(a, b []int) []int {
	for _, biss := range b {
		a = append(a, biss)
	}

	return a
}
*/
