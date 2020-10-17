from typing import List

class Solution:
    
    # Runtime: 36 ms, faster than 81.02% of Python3 online submissions
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        pt1 = m - 1
        pt2 = n - 1
        pt = m + n -1
        
        while pt2 >= 0:
            if pt1 < 0 or nums1[pt1] < nums2[pt2]:
                nums1[pt] = nums2[pt2]
                pt2 -= 1
            else:
                nums1[pt] = nums1[pt1]
                pt1 -= 1
            pt -= 1


    # Runtime: 36 ms, faster than 81.02% of Python3 online submissions
    def merge_v2(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        nums1[:] = sorted(nums1[:m] + nums2)