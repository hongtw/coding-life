from typing import List

class Solution:

    # Runtime: 28 ms, faster than 94.50% of Python3 online submissions for Sort Colors.
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        left = 0
        point = left
        right = len(nums) - 1
        
        while point <= right and left < right:
            
            while left < right and nums[left] == 0:
                left += 1
                point = left
            while left < right and nums[right] == 2:
                right -= 1

            if nums[point] == 0:
                nums[left], nums[point] = nums[point], nums[left]
            elif nums[point] == 1:
                point += 1
            else:
                nums[point], nums[right] = nums[right], nums[point]