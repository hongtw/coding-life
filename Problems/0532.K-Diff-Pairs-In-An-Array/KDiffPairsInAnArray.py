from typing import List
from collections import Counter

class Solution:

    def findPairs(self, nums: List[int], k: int) -> int:
        if k < 0:
            return 0

        res = 0
        nums = Counter(nums)
        if k == 0:
            for n in nums:
                if nums[n] > 1:
                    res += 1
            return res

        for n in nums:
            if n + k in nums:
                res += 1
        return res