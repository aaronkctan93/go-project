package main

import "fmt"

func main() {

	// nums := []int{5, 7, 1, 10, 2}
	// Q2 : [4,1,5,9]
	// Q2 : [1,3,2,6]
	students := []int{1, 3, 2, 6}
	seats := []int{4, 1, 5, 9}

	fmt.Println(minMovesToSeat(seats, students))
}

func minMovesToSeat(seats []int, students []int) int {
	countMove := 0
	checkIsSmallest := 0
	recordPosition := 0
	distance := make([]int, len(seats))

	for _, student := range students {
		for k, seat := range seats {
			if seat != -1 {

				distance[k] = student - seat
				if distance[k] < 0 {
					distance[k] *= -1
				}

				if checkIsSmallest == 0 {
					checkIsSmallest = distance[k]
					recordPosition = k
				} else {
					if checkIsSmallest > distance[k] {
						checkIsSmallest = distance[k]
						recordPosition = k
					} else if checkIsSmallest == distance[k] {
						// fmt.Println("SAME")
						if recordPosition < k {
							recordPosition = k
						}
					}

				}

			}
		}
		seats[recordPosition] = -1
		fmt.Println(checkIsSmallest)
		countMove += checkIsSmallest
		checkIsSmallest = 0

	}

	return countMove
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
