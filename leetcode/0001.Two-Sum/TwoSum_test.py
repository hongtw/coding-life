from TwoSum import Solution
import pytest

SOL = Solution()

@pytest.mark.parametrize(
    "nums, target, ans", 
    [
        ([2, 7, 11, 15], 9, [0, 1]), 
        ([4, 6, 2, 1], 8, [1, 2]),
    ]
)
def test(nums, target, ans):
    assert SOL.twoSum(nums, target) == ans