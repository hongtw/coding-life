package leetcode

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Fibonacci Number.
func fib(N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	prev := 1
	curr := 1
	count := 0
	for i := 2; i < N; i++ {
		count = prev + curr
		prev = curr
		curr = count
	}
	return count
}
