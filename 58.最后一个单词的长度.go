package main

func lengthOfLastWord(s string) int {
	flag := false
	resultLength := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			flag = true
		}
		if s[i] != ' ' {
			resultLength++
		} else if flag {
			return resultLength
		}
	}
	return resultLength
}
