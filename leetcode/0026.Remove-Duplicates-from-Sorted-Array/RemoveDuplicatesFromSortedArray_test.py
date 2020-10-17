from RemoveDuplicatesFromSortedArray import Solution
import pytest

SOL = Solution()

TEST_SUITS = [
        ([1,1,2], 2, [1, 2]), 
        ([0,0,1,1,1,2,2,3,3,4], 5, [0, 1, 2, 3, 4]), 
        ([1, 1], 1, [1]),
        ([], 0, [])
]
@pytest.mark.parametrize(
    "nums, ans, check_nums", 
    TEST_SUITS
)
def test_v1(nums, ans, check_nums):
    res = SOL.removeDuplicates(nums)
    assert  res == ans, f"Got ans={res}"
    for i, n in enumerate(check_nums):
        assert nums[i] == n

@pytest.mark.parametrize(
    "nums, ans, check_nums", 
    TEST_SUITS
)
def test_v2(nums, ans, check_nums):
    res = SOL.removeDuplicates_v2(nums)
    assert  res == ans, f"Got ans={res}"
    for i, n in enumerate(check_nums):
        assert nums[i] == n