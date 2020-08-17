
from MergeSortedArray import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ("Example 1", [1,2,3,0,0,0], 3, [2,5,6], 3, [1,2,2,3,5,6]),
    ("Example 2", [0], 0, [1], 1, [1]),
    ("Example 3", [2, 0], 1, [1], 1, [1, 2]),
]

@pytest.mark.parametrize(
    "example, nums1, m, nums2, n, ans", 
    deepcopy(TEST_SUITS)
)
def test(example, nums1, m, nums2, n, ans):
    SOL.merge(nums1, m, nums2, n)
    assert  nums1 == ans

@pytest.mark.parametrize(
    "example, nums1, m, nums2, n, ans", 
    deepcopy(TEST_SUITS)
)
def test_v2(example, nums1, m, nums2, n, ans):
    SOL.merge_v2(nums1, m, nums2, n)
    assert  nums1 == ans