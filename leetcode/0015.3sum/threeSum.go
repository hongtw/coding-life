package leetcode

import "sort"

// Runtime: 84 ms, faster than 35.65% of Go online submissions
func threeSum(nums []int) [][]int {
	counter := make(map[int]int)
	res := make([][]int, 0)

	for _, n := range nums {
		counter[n]++
	}
	uniqNums := []int{}
	for k := range counter {
		uniqNums = append(uniqNums, k)
	}
	sort.Ints(uniqNums)

	for i := range uniqNums {
		for j := i + 1; j < len(uniqNums); j++ {
			if uniqNums[i]*2+uniqNums[j] == 0 && counter[uniqNums[i]] >= 2 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			} else if uniqNums[i]+uniqNums[j]*2 == 0 && counter[uniqNums[j]] >= 2 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			} else {
				r := 0 - uniqNums[i] - uniqNums[j]
				if counter[r] > 0 && r > uniqNums[j] {
					res = append(res, []int{uniqNums[i], uniqNums[j], r})
				}
			}

		}
	}
	if counter[0] >= 3 {
		res = append(res, []int{0, 0, 0})
	}
	return res
}

// Runtime: 28 ms, faster than 94.38% of Go online submissions
func threeSumV2(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	length := len(nums)

	for i := 0; i < length; i++ {
		for j, k := i+1, length-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j+1 < length && j < k && nums[j] == nums[j+1] {
					j++
				}
				j++
			} else {
				j++
			}
		}
		for i+1 < length && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
