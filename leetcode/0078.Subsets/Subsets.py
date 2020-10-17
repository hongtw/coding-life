from typing import List
from itertools import combinations

class Solution:

    # Runtime: 28 ms, faster than 96.15% of Python3 online submissions for Subsets
    def subsets(self, nums: List[int]) -> List[List[int]]:
        return [
            list(tup)
            for i in range(len(nums)+1)
            for tup in combinations(nums, i)
        ]
    
    # Runtime: 24 ms, faster than 99.21% of Python3 online submissions for Subsets.
    def subsetsV2(self, nums: List[int]) -> List[List[int]]:
        res = [[]]
        for n in nums:
            res += [tup + [n] for tup in res ]
            print("****", res)
        return res