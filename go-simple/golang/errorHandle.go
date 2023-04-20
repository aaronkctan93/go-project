package main

import (
	"errors"
	"fmt"
)

func f1(nums int) (int, error) {
	if nums < 0 {
		return -1, errors.New("參數錯誤")
	}

	return 2 * nums, nil
}

type arrError struct {
	arg     int
	message string
}

func (e *arrError) Error() string {
	return fmt.Sprintf("%d -> %s", e.arg, e.message)
}

func f2(nums int) (int, error) {
	if nums < 0 {
		return -1, &arrError{nums, "參數不為負數"}
	}

	return 2 * nums, nil
}

func main() {

	fmt.Println(f1(2))
	fmt.Println(f1(-1))

	fmt.Println(f2(8))
	fmt.Println(f2(-5))

}
