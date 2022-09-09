package main

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	prevAns := countAndSay(n - 1)
	char := prevAns[0:1]
	prevChar := char
	count := 1
	ans := ""
	if len(prevAns) > 1 {
		for _, ch := range prevAns[1:] {
			char = string(ch)
			if char != prevChar {
				ans += strconv.Itoa(count) + prevChar
				prevChar = char
				count = 0
			}
			count += 1
		}
	}
	ans += strconv.Itoa(count) + char
	return ans
}

// def countAndSay(self, n: int) -> str:
// if n == 1:
// 	return "1"

// prevAns = self.countAndSay(n - 1)
// char = prevChar = prevAns[0]
// count = 1
// ans = ""
// for char in prevAns[1:]:
// 	if char != prevChar:
// 		ans += f"{count}{prevChar}"
// 		prevChar = char
// 		count = 0
// 	count += 1

// ans += f"{count}{char}"

// return ans
