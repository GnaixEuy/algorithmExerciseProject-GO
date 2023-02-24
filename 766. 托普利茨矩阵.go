package main

func isToeplitzMatrix(matrix [][]int) bool {
	colNum := len(matrix[0])
	rowNum := len(matrix)
	for i := 1; i < colNum; i++ {
		for j := 1; j < rowNum; j++ {
			if matrix[j-1][i-1] != matrix[j][i] {
				return false
			}
		}
	}
	return true
}
