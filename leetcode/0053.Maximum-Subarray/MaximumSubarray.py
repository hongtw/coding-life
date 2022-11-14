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

    # https://leetcode.com/problems/maximum-subarray/discuss/1595195/C%2B%2BPython-7-Simple-Solutions-w-Explanation-or-Brute-Force-%2B-DP-%2B-Kadane-%2B-Divide-and-Conquer
    def maxSubArrayV2(self, nums: List[int]) -> int:
        # Time Complexity : O(N), for iterating over nums once
        # Space Complexity : O(1), only constant extra space is being used.
        cur_max = 0
        res_max = 0
        for n in nums:
            cur_max = max(n, cur_max+n)
            res_max = max(res_max, cur_max)
        return res_max
            
    def maxSubArrayV3(self, nums: List[int]) -> int:
        # Time Complexity : O(N), for iterating over nums once
        # Space Complexity : O(1)
        for i in range(1, len(nums)):
            nums[i] = max(nums[i], nums[i-1]+nums[i])
        return max(nums)
            