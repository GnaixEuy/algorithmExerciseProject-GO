package main

func reorganizeString(s string) string {
	chNum := make([]int, 26)
	sum := len(s)
	max := 0
	maxI := 0
	for _, ch := range s {
		chNum[ch-'a']++
		if chNum[ch-'a'] > max {
			max = chNum[ch-'a']
			maxI = int(ch - 'a')
		}
	}
	if max > (1+sum)/2 {
		return ""
	}
	ji := 0
	ou := 1
	str := make([]byte, sum)
	for ji = 0; chNum[maxI] > 0; ji += 2 {
		str[ji] = byte(maxI + 'a')
		chNum[maxI]--
	}
	for i := range chNum {
		if chNum[i] == 0 {
			continue
		}
		for ; chNum[i] > 0; chNum[i]-- {
			if ji < sum {
				str[ji] = byte(i + 'a')
				ji += 2
			} else {
				str[ou] = byte(i + 'a')
				ou += 2
			}
		}
	}
	return string(str)
}
