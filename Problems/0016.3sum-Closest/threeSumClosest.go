package leetcode

import "sort"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for 3Sum Closest.
func threeSumClosest(nums []int, target int) int {
	diff := 10000
	res := 0
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length-2; i++ {
		for j, k := i+1, length-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}
			if abs(sum-target) < diff {
				diff, res = abs(sum-target), sum
			}
			if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return res
}
