package main

func runningSum(nums []int) []int {
	recursion1480(nums, 0)
	return nums
}

func recursion1480(nums []int, i int) {
	if i == len(nums) {
		return
	}
	if i != 0 {
		nums[i] += nums[i-1]
	}
	recursion1480(nums, i+1)
}
