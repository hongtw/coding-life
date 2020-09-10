from typing import List

class Solution:

    # Runtime: 40 ms, faster than 97.22% of Python3 online submissions for Combination Sum II.
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort()
        def combinHelper(candidates, target, prefixRes, res):
            for i, n in enumerate(candidates):
                if i > 0 and n == candidates[i-1]:
                    continue
                if n == target:
                    res.append(prefixRes + [n])
                elif n < target:
                    combinHelper(candidates[i+1:], target - n, prefixRes + [n], res)
                else:
                    break
        res = []
        combinHelper(candidates, target, [], res)
        return res
