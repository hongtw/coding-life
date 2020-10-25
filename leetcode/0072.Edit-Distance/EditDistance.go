package main

import "fmt"

func mins(value int, values ...int) int {
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

func levRecur(word1 string, word2 string, memo map[[2]string]int) int {

	if l1, l2 := len(word1), len(word2); l1 == 0 || l2 == 0 {
		if l1 > l2 {
			return l1
		}
		return l2
	}
	if val, ok := memo[[2]string{word1, word2}]; ok {
		return val
	}
	if word1[0] == word2[0] {
		return levRecur(word1[1:], word2[1:], memo)
	}
	dist := 1 + mins(
		levRecur(word1[1:], word2[1:], memo),
		levRecur(word1[1:], word2, memo),
		levRecur(word1, word2[1:], memo),
	)
	memo[[2]string{word1, word2}] = dist
	return dist
}

// Runtime: 40 ms, faster than 10.57% of Go online submissions for Edit Distance.
func minDistance(word1 string, word2 string) int {
	memo := make(map[[2]string]int)
	return levRecur(word1, word2, memo)
}

func popQueue(queue *[][3]int) [3]int {
	tuple := (*queue)[0]
	*queue = (*queue)[1:]
	return tuple
}

func pushQueue(queue *[][3]int, tuples ...[3]int) {
	*queue = append(*queue, tuples...)
}

// Runtime: 8 ms, faster than 43.90% of Go online submissions for Edit Distance.
func minDistanceQueue(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	queue := [][3]int{{0, 0, 0}}
	isVisited := make(map[[2]int]bool)
	for {
		tuple := popQueue(&queue)
		p1, p2, dist := tuple[0], tuple[1], tuple[2]
		if _, ok := isVisited[[2]int{p1, p2}]; ok {
			continue
		}
		isVisited[[2]int{p1, p2}] = true
		for p1 < l1 && p2 < l2 && word1[p1] == word2[p2] {
			p1++
			p2++
		}
		if p1 >= l1 && p2 >= l2 {
			return dist
		}
		dist++
		pushQueue(
			&queue,
			[3]int{p1 + 1, p2 + 1, dist},
			[3]int{p1, p2 + 1, dist},
			[3]int{p1 + 1, p2, dist},
		)
	}
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Edit Distance.
func minDistanceDP(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)

	// creat DP Table
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}
	for i := range dp[0] {
		dp[0][i] = i
	}
	for i := range dp {
		dp[i][0] = i
	}

	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + mins(
					dp[i-1][j-1],
					dp[i-1][j],
					dp[i][j-1],
				)
			}
		}
	}
	return dp[l1][l2]
}

func main() {
	fmt.Println(minDistance("kitten", "sitting"))
	fmt.Println(minDistanceQueue("kitten", "sitting"))
	fmt.Println(minDistanceDP("kitten", "sitting"))
}
