package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hm := make(map[rune]rune)
	hmT := make(map[rune]rune)
	for k, v := range s {
		if val, ok := hm[v]; ok {
			if rune(t[k]) != val {
				return false
			}
		} else {
			hm[v] = rune(t[k])
		}
		if val, ok := hmT[rune(t[k])]; ok {
			if val != v {
				return false
			}
		} else {
			hmT[rune(t[k])] = v
		}
	}
	return true
}
