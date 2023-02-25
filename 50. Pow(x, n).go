package main

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / quickMul(x, -n)
	} else {
		return quickMul(x, n)
	}
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	ans := quickMul(x, n/2)
	if n%2 == 0 {
		return ans * ans
	} else {
		return ans * ans * x
	}
}
