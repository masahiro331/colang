package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(rotate(slice, 3))
}

func rotate(nums []int, r int) []int {
	nums = append(nums[len(nums)-r%len(nums):], nums[:len(nums)-r%len(nums)]...)
	return nums
}
