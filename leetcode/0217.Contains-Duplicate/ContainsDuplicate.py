from typing import List

class Solution:

    # Runtime: 120 ms, faster than 95.29% of Python3 online submissions
    def containsDuplicate(self, nums: List[int]) -> bool:
        rec = set()
        for n in nums:
            if n in rec:
                return True
            rec.add(n)
        return False

    # Runtime: 120 ms, faster than 95.29% of Python3 online submissions
    def containsDuplicate_v2(self, nums: List[int]) -> bool:
        return len(set(nums)) != len(nums)