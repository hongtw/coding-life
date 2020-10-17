from typing import List

class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        if k == 0:
            return False
        k_cache = set()
        for i in range(len(nums)):
            if nums[i] in k_cache:
                return True
            if i >= k:
                k_cache.remove(nums[i-k])
            k_cache.add(nums[i])
        return False

    # Runtime: 88 ms, faster than 95.61% of Python3 online submissions for Contains Duplicate II.
    def containsNearbyDuplicate_v2(self, nums: List[int], k: int) -> bool:
        if k == 0:
            return False
        k_cache = {}
        for i, num in enumerate(nums):
            if num in k_cache and i - k_cache[num] <= k:
                return True
            k_cache[num] = i
        return False