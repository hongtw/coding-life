from typing import List

class Solution:

    # Runtime: 96 ms, faster than 96.03% of Python3 online submissions for Reshape the Matrix.
    def matrixReshape(self, nums: List[List[int]], r: int, c: int) -> List[List[int]]:
        h = len(nums)
        w = len(nums[0])
        if r * c != h * w: return nums

        res = [[None] * c for i in range(r)]
        for i in range(h*w):
            res[i//c][i%c] = nums[i//w][i%w]

        return res