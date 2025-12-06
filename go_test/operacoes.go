package main

import (
	"fmt"
)

func main() {
	x := Sum(1, 2, 3)
	y := Multiply(10, 10)
	w := Subtract(5, 10)
	z, err := Divide(20)
	if err != nil {
		fmt.Println("Division error:", err)
		return
	}

	fmt.Println(x, y, w, z)
}

func Sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func Subtract(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	total := nums[0]
	for _, v := range nums[1:] {
		total -= v
	}
	return total
}

func Multiply(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	total := 1
	for _, v := range nums {
		total *= v
	}
	return total
}

func Divide(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("no parameters provided")
	}

	total := nums[0]
	for _, v := range nums[1:] {
		if v == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		total /= v
	}

	return total, nil
}
