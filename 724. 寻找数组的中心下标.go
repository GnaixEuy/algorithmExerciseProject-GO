package main

func pivotIndex(nums []int) int {
	sumNum := 0
	for _, num := range nums {
		sumNum += num
	}
	tmpNum := 0
	for i, num := range nums {
		if tmpNum<<1+num == sumNum {
			return i
		}
		tmpNum += num
	}
	return -1
}
