package leetcode

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Sort Colors.
func sortColors(nums []int)  {
	left := 0
	point := left
	right := len(nums) - 1
	for point <= right && left < right {
		for left < right && nums[left] == 0 {
			left++
			point = left
		}
		for left < right && nums[right] == 2 {
			right--
		}
		if nums[point] == 0 {
			nums[left], nums[point] = nums[point], nums[left]
		} else if nums[point] == 1 {
			point++
		} else{
			nums[right], nums[point] = nums[point], nums[right]
		}

	}
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Sort Colors.
func sortColorsV2(nums []int)  {
	p0 := 0
	p1 := 0
	p2 := 0
	for _, num := range nums {
		fmt.Println(p0, p1, p2)
		if num == 0 {
			nums[p2] = 2
			p2++
			nums[p1] = 1
			p1++
			nums[p0] = 0
			p0++
		} else if num == 1 {
			nums[p2] = 2
			p2++
			nums[p1] = 1
			p1++
		} else {
			p2++
		}
	}
}