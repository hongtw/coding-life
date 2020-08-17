from typing import List

# Runtime: 60 ms, faster than 96.88% of Python3 online submissions
class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        maxsum = None
        subsum = 0
        for n in nums:
            subsum += n
            if maxsum is None or subsum > maxsum:
                maxsum = subsum
            if subsum < 0:
                subsum = 0
        return maxsum