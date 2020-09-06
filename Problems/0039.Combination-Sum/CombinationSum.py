from typing import List, Set

class Solution:

    # Runtime: 56 ms, faster than 87.55% of Python3 online submissions for Combination Sum.
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        print("XXX ",candidates, target )
        res = []
        for i, n in enumerate(candidates):
            if n == target:
                res.append([n])
            elif n < target:
                subRes = self.combinationSum(candidates[i:], target - n)
                res.extend([[n] + r for r in subRes])


        print("====", candidates, target, res)
        return res

    # Runtime: 56 ms, faster than 87.55% of Python3 online submissions for Combination Sum.
    def combinationSumV2(self, candidates: List[int], target: int) -> List[List[int]]:
        def combinHelper(candidates, target, prefixRes):
            for i, n in enumerate(candidates):
                if n == target:
                    res.append(prefixRes + [n])
                elif n < target:
                    combinHelper(candidates[i:], target - n, prefixRes + [n])
        res = []
        combinHelper(candidates, target, [])
        return res

    # Runtime: 48 ms, faster than 96.82% of Python3 online submissions for Combination Sum.
    def combinationSumV3(self, candidates: List[int], target: int) -> List[List[int]]:
        def combinHelper(candidates, target, prefixRes, res):
            for i, n in enumerate(candidates):
                if n == target:
                    res.append(prefixRes + [n])
                elif n < target:
                    combinHelper(candidates[i:], target - n, prefixRes + [n], res)
        res = []
        combinHelper(candidates, target, [], res)
        return res