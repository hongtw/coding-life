package leetcode

func twoSum(nums []int, target int) []int {
	rec := make(map[int]int)
	for idx, n := range nums {
		m := target - n
		if loc, isExist := rec[m]; isExist {
			return []int{loc, idx}
		}
		rec[n] = idx
	}
	return []int{}
}
