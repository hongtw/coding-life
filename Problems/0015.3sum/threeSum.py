from typing import List, Tuple

class Solution:

    # Runtime: 360 ms, faster than 98.51% of Python3 online submissions for 3Sum.
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        res = []
        pos_counter = {}
        neg_counter = {}
        for n in nums:
            if n >= 0:
                pos_counter[n] = pos_counter.get(n, 0) + 1
            else:
                neg_counter[n] = neg_counter.get(n, 0) + 1


        for pos in pos_counter:
            for neg in neg_counter:
                x = 0 - pos - neg
                if x in pos_counter or x in neg_counter:

                    if (x == pos and pos_counter[x] > 1) or \
                        (x == neg and neg_counter[x] > 1):
                        res.append([neg, x, pos])
                    elif x > pos:
                        res.append([neg, pos, x])
                    elif x < neg:
                        res.append([x, neg, pos])

        if pos_counter.get(0, 0) >= 3:
            res.append([0, 0, 0])
        return res

    def threeSumV2(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        counter = {}
        for n in nums:
            counter[n] = counter.get(n, 0) + 1

        length = len(nums)
        for i in range(length - 2):
            fst = nums[i]
            sec = nums[i + 1]
            
            for j in range(i+2, length-1):
                nums[i]