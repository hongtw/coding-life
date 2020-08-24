from typing import List
from collections import Counter

class Solution:

    # Runtime: 128 ms, faster than 88.59% of Python3 online submissions for K-diff Pairs in an Array.
    def findPairs(self, nums: List[int], k: int) -> int:
        if k < 0:
            return 0
        rec = set()
        res = set()
        for n in nums:
            if n + k in rec:
                if n + k > n:
                    res.add(f"{n},{n+k}")
                else:
                    res.add(f"{n+k},{n}")

            if n - k in rec:
                if n - k > n:
                    res.add(f"{n},{n - k}")
                else:
                    res.add(f"{n - k},{n}")

            rec.add(n)
        return len(res)

    # Runtime: 112 ms, faster than 99.81% of Python3 online submissions
    def findPairsV2(self, nums: List[int], k: int) -> int:
        if k < 0:
            return 0

        res = 0
        if k == 0:
            nums = Counter(nums)
            for n in nums:
                if nums[n] > 1:
                    res += 1
            return res

        nums = set(nums)
        sorted_nums = sorted(nums)

        for n in sorted_nums:
            if n + k in nums:
                res += 1
        return res