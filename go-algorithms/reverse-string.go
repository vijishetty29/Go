package main

func Reverse(s string) string {
	strs := []rune(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}
	return string(strs)
}
