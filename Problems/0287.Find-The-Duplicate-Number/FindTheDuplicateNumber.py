from typing import List

class Solution:

    # Runtime: 52 ms, faster than 99.91% of Python3 online submissions for Find the Duplicate Number.
    def findDuplicate(self, nums: List[int]) -> int:
        cache = set()
        for n in nums:
            if n in cache:
                return n
            cache.add(n)
        return 1

    # Runtime: 52 ms, faster than 99.91% of Python3 online submissions for Find the Duplicate Number.
    def findDuplicateV2(self, nums: List[int]) -> int:
        pt = 0
        while True:
            if nums[pt] == 0:
                return pt
            newpt = nums[pt]
            nums[pt] = 0
            pt = newpt
            
            # Can't do this in python, pt will be assigned first
            # pt, nums[pt] = nums[pt], 0