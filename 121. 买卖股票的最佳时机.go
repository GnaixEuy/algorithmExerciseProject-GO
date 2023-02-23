package main

import "math"

func maxProfit(prices []int) int {
	result := 0.0
	getLow := math.MaxFloat64
	for i := range prices {
		getLow = math.Min(getLow, float64(prices[i]))
		result = math.Max(result, float64(prices[i])-getLow)
	}
	return int(result)
}
