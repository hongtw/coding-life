from typing import List

class Solution:

    # Runtime: 32 ms, faster than 90.29% of Python3 online submissions for First Missing Positive.
    def firstMissingPositive(self, nums: List[int]) -> int:
        nums = set(nums)
        for i in range(1, len(nums) + 1):
            if i not in nums:
                return i

        return len(nums) + 1