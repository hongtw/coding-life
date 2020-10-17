package main

import (
	"fmt"
	"strings"
)

func searchSubstr(fullText string, searchText string, allowOverlap bool) int {
	if fullText == "" || searchText == "" {
		return 0
	}
	startPos := 0
	count := 0
	step := 1
	if !allowOverlap {
		step = len(searchText)
	}
	for matchedIndex := 0; matchedIndex > -1; {
		matchedIndex = strings.Index(fullText[startPos:], searchText)
		if matchedIndex > -1 {
			count++
			startPos += (matchedIndex + step)
		}
	}
	return count
}

func main() {
	fullText := "aa_bb_cc_dd_bb_e"
	searchText := "bb"
	allowOverlap := true
	cnt := searchSubstr(fullText, searchText, allowOverlap)
	fmt.Println(fullText, "->", searchText, "->", cnt)
}
