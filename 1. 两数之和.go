package main

func twoSum(nums []int, target int) []int {
	hmap := make(map[int]int)
	for k, v := range nums {
		if preIndex, ok := hmap[target-v]; ok {
			return []int{preIndex, k}
		} else {
			hmap[v] = k
		}
	}
	return []int{}
}
