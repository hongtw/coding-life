from typing import List

class Solution:

    # Runtime: 44 ms, faster than 95.39% of Python3 online submissions for Move Zeroes.
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        pt1 = 0
        pt2 = 1
        length = len(nums)
        while True:
            while pt1 < length and nums[pt1] != 0:
                pt1 += 1
            while pt2 < length and (pt2 < pt1 or nums[pt2] == 0):
                pt2 += 1
            if pt1 == length or pt2 == length:
                break
            nums[pt1], nums[pt2] = nums[pt2], nums[pt1]