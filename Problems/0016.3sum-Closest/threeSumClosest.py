from typing import List
from collections import Counter

class Solution:

    # Runtime: 328 ms, faster than 23.08% of Python3 online submissions for 3Sum Closest.
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        closest = 10000
        output = 0
        counter = Counter(nums)

        uniqNums = list(counter)
        uniqlen = len(uniqNums)
        for i in range(uniqlen):
            print("closest", closest, output)
            fst = uniqNums[i]
            if counter[fst] > 2 and abs(fst * 3 - target) < closest:
                closest = abs(fst * 3 - target)
                output = fst * 3
            
            print("xxx", fst)
            for j in range(i+1, uniqlen):

                sec = uniqNums[j]
                print("OOO", fst, sec)
                print("====", abs(fst*2 + sec - target))
                if counter[fst] > 1 and abs(fst*2 + sec - target) < closest:
                    closest = abs(fst*2 + sec - target)
                    output = fst*2 + sec
                if counter[sec] > 1 and abs(fst + sec*2 - target) < closest:
                    closest = abs(fst + sec*2 - target)
                    output = fst + sec*2

                for k in range(j+1, uniqlen):
                    trd = uniqNums[k]

                    diff = target - fst - sec - trd
                    if abs(diff) < closest:
                        closest = abs(diff)
                        output = fst + sec + trd 

                    if output == target:
                        return output
        return output

    # Runtime: 116 ms, faster than 90.56% of Python3 online submissions for 3Sum Closest.
    def threeSumClosestV2(self, nums: List[int], target: int) -> int:
        res = 0
        diff = 10000
        nums.sort()
        length = len(nums)
        for i in range(length-2):
            j = i + 1
            k = length - 1
            while j < k:
                sum_ = nums[i] + nums[j] + nums[k]
                if sum_ == target:
                    return target 
                tmp_diff = abs(sum_ - target)
                if tmp_diff < diff:
                    diff, res = tmp_diff, sum_
                if sum_ > target:
                    k -= 1
                else:
                    j += 1
        return res