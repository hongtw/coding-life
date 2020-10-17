from typing import List

class Solution:

    # Runtime: 28 ms, faster than 93.19% of Python3 online submissions
    def removeElement(self, nums: List[int], val: int) -> int:
        recorder = 0
        for index in range(len(nums)):
            if nums[index] != val:
                nums[recorder] = nums[index]
                recorder += 1
        return recorder
    
    # Runtime: 28 ms, faster than 93.19% of Python3 online submissions
    def removeElement_v2(self, nums: List[int], val: int) -> int:
        length = len(nums)
        index = 0
        while index < length:
            if nums[index] == val:
                del nums[index]
                length -= 1
            else:
                index += 1
        return index

