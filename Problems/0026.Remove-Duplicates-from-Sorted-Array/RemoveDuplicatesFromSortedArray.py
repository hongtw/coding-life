from typing import List

class Solution:

    # Runtime: 76 ms, faster than 98.15% of Python3 online submissions
    # Memory Usage: 15.7 MB, less than 66.66% of Python3 online submissions
    def removeDuplicates(self, nums: List[int]) -> int:
        total = len(nums) 
        if total <= 1:
            return total

        recorder = 0
        for finder in range(total):
            if nums[finder] != nums[recorder]:
                recorder += 1
                nums[recorder] = nums[finder]
        return recorder + 1

    # Runtime: 80 ms, faster than 94.29% of Python3 online submissions
    def removeDuplicates_v2(self, nums: List[int]) -> int:
        nums[:] = sorted(set(nums))
        return len(nums)