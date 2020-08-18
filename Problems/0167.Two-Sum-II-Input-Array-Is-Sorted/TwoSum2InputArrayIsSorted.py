from typing import List

class Solution:

    # Runtime: 56 ms, faster than 98.27% of Python3 online submissions
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        rec = {}
        for i, n in enumerate(numbers):
            if n in rec:
                return [rec[n] + 1, i + 1]
            else:
                key = target - n
                rec[key] = i
        return []