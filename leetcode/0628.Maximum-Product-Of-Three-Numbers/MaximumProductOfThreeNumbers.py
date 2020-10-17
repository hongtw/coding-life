from typing import List

class Solution:
    # Runtime: 276 ms, faster than 76.58% of Python3 online submissions for Maximum Product of Three Numbers.
    def maximumProduct(self, nums: List[int]) -> int:
        top1 = -1000
        top2 = -1000
        top3 = -1000
        neg1 = 0
        neg2 = 0
        for n in nums:
            if n > top1:
                # top3, top2, top1 = top2, top1, n
                top1, top2, top3 = n, top1, top2
            elif n > top2:
                top3, top2 = top2, n
            elif n > top3:
                top3 = n
                
            if n < neg1:
                neg2, neg1 = neg1, n
            elif n < neg2:
                neg2 = n

        print("=========", nums, "===", top1, top2, top3, neg1, neg2)
        return max(top1 * top2 * top3, top1 * neg1 * neg2)

    # Runtime: 268 ms, faster than 92.32% of Python3 online submissions for Maximum Product of Three Numbers.
    def maximumProductV2(self, nums: List[int]) -> int:
        nums.sort()
        return max(nums[-1]*nums[0]*nums[1],  nums[-1]*nums[-2]*nums[-3])
