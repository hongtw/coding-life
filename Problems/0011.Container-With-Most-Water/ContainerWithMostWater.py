from typing import List

class Solution:

    # Runtime: 112 ms, faster than 99.74% of Python3 online submissions for Container With Most Water.
    def maxArea(self, height: List[int]) -> int:
        l = 0
        r = len(height) - 1
        max_area = 0
        while l < r:
            if height[l] < height[r]:
                max_area = max(max_area, height[l] * (r-l))
                l += 1
            else:
                max_area = max(max_area, height[r] * (r-l))
                r -= 1
        return max_area
