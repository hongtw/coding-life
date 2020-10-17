package leetcode

// Runtime: 12 ms, faster than 97.87% of Go online submissions
// Memory Usage: 6 MB, less than 80.14% of Go online submissions
func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		return false
	}
	kCache := make(map[int]bool, k)
	for i := range nums {
		if _, isExist := kCache[nums[i]]; isExist {
			return true
		}
		if i >= k {
			delete(kCache, nums[i-k])
		}
		kCache[nums[i]] = true
	}
	return false
}

func containsNearbyDuplicateV2(nums []int, k int) bool {
	if k == 0 {
		return false
	}
	kCache := make(map[int]int, len(nums))
	for i, num := range nums {
		if loc, isExist := kCache[num]; isExist {
			if i-loc <= k {
				return true
			}
		}
		kCache[num] = i
	}
	return false
}
