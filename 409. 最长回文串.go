package main

func longestPalindrome(s string) int {
	if len(s) == 1 {
		return 1
	}
	dui := 0
	dan := 0
	hmap := make(map[int32]int)
	for _, ch := range s {
		hmap[ch]++
		if hmap[ch] != 0 && hmap[ch]%2 == 0 {
			dui++
			dan--
		} else if hmap[ch] != 0 {
			dan++
		}
	}
	if dan > 0 {
		return dui*2 + 1
	}
	return dui * 2
}
