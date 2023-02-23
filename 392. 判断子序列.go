package main

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	i := 0
	for j := range t {
		if s[i] == t[j] {
			if i == len(s)-1 {
				return true
			}
			i++
		}
	}
	return false
}
