from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        rec = {}
        for idx, n in enumerate(nums):
            m = target - n
            if m in rec:
                return [rec[m], idx]
            rec[n] = idx
